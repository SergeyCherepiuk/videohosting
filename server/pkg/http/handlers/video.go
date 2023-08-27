package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/SergeyCherepiuk/videohosting/pkg/bucket"
	"github.com/SergeyCherepiuk/videohosting/pkg/internal/reader"
	"github.com/SergeyCherepiuk/videohosting/pkg/internal/sse"
	"github.com/labstack/echo/v4"
)

type VideoHandler struct {
	bucket bucket.BucketService
}

func NewVideoHandler(bucket bucket.BucketService) *VideoHandler {
	return &VideoHandler{bucket: bucket}
}

func (handler VideoHandler) Upload(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return err
	}

	eventId := 0
	reader := reader.NewReaderWithProgress(file, func(bytesRead int) {
		event := sse.Event{
			ID:        eventId,
			EventName: "progress",
			Data:      fmt.Sprintf("%.2f", float32(bytesRead)/float32(fileHeader.Size)),
		}	
		if err := event.Send(c.Response()); err != nil {
			c.Logger().Errorf("failed to send a SSE: %w", err)
		}
		eventId++
	})

	if err := handler.bucket.Upload(context.Background(), fileHeader.Filename, reader); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
