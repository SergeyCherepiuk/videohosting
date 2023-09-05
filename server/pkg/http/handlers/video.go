package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/SergeyCherepiuk/videohosting/domain"
	"github.com/SergeyCherepiuk/videohosting/pkg/http/validation"
	"github.com/SergeyCherepiuk/videohosting/pkg/internal/reader"
	"github.com/SergeyCherepiuk/videohosting/pkg/internal/sse"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type VideoHandler struct {
	videoService domain.VideoService
	bucket       domain.BucketService
}

func NewVideoHandler(
	videoService domain.VideoService,
	bucket domain.BucketService,
) *VideoHandler {
	return &VideoHandler{
		videoService: videoService,
		bucket:       bucket,
	}
}

func (handler VideoHandler) GetByUUID(c echo.Context) error {
	uuid, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	video, err := handler.videoService.GetByUUID(uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, video)
}

func (handler VideoHandler) Upload(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")

	// Retrieving data from multipart form
	previewFileHeader, err := c.FormFile("preview")
	if previewFileHeader == nil || err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	previewFile, err := previewFileHeader.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	videoFileHeader, err := c.FormFile("video")
	if videoFileHeader == nil || err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	videoFile, err := videoFileHeader.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	request := validation.UploadVideoRequest{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Preview:     previewFileHeader,
		Video:       videoFileHeader,
	}

	// Validating the data
	if err := request.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	video := domain.Video{
		ID:          uuid.New(),
		AuthorID:    uuid.MustParse("657e846b-073b-454a-999b-2ce6157452cc"), // hardcoded for now (no auth)
		Title:       request.Title,
		Description: request.Description,
		PreviewID:   uuid.New(),
		ViewsCount:  0,
		CreatedAt:   time.Now(),
	}

	// Storing video details in database
	if err := handler.videoService.Upload(video); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Creating the reader that will send events (SSE) with progress to the client
	eventId := 0
	reader := reader.NewReaderWithProgress(videoFile, func(bytesRead int) {
		event := sse.Event{
			ID:        eventId,
			EventName: "progress",
			Data: fmt.Sprintf(
				`{ "progress": %.2f }`,
				float32(bytesRead)/float32(videoFileHeader.Size),
			),
		}
		if err := event.Send(c.Response()); err != nil {
			c.Logger().Errorf("failed to send a SSE: %w", err)
		}
		eventId++
	})

	// Uploading preview to the bucket
	previewFilePath := fmt.Sprintf("%s/%s", os.Getenv("S3_PREVIEW_FOLDER"), video.PreviewID.String())
	if err := handler.bucket.Upload(
		context.Background(), previewFilePath, previewFileHeader.Header.Get("Content-Type"), previewFile,
	); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Uploading video to the bucket
	videoFilePath := fmt.Sprintf("%s/%s", os.Getenv("S3_VIDEO_FOLDER"), video.ID.String())
	if err := handler.bucket.Upload(
		context.Background(), videoFilePath, videoFileHeader.Header.Get("Content-Type"), reader,
	); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}
