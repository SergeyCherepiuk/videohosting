package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/SergeyCherepiuk/videohosting/pkg/bucket"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PreviewHandler struct {
	bucket bucket.BucketService
}

func NewPreviewHandler(bucket bucket.BucketService) *PreviewHandler {
	return &PreviewHandler{bucket: bucket}
}

func (handler PreviewHandler) GetByUUID(c echo.Context) error {
	uuid, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	preview, mime, err := handler.bucket.Get(
		context.Background(),
		fmt.Sprintf("%s/%s", os.Getenv("S3_PREVIEW_FOLDER"), uuid.String()),
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.Blob(http.StatusOK, mime, preview)
}
