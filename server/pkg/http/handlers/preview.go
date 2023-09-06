package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/SergeyCherepiuk/videohosting/domain"
	"github.com/SergeyCherepiuk/videohosting/pkg/database/redis"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PreviewHandler struct {
	bucket domain.BucketService
}

func NewPreviewHandler(bucket domain.BucketService) *PreviewHandler {
	return &PreviewHandler{bucket: bucket}
}

func (handler PreviewHandler) GetByUUID(c echo.Context) error {
	uuid, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	ctxWithTtl := context.WithValue(context.Background(), redis.CtxTtlKey, 24 * time.Hour)
	preview, mime, err := handler.bucket.Get(
		ctxWithTtl, fmt.Sprintf("%s/%s", os.Getenv("S3_PREVIEW_FOLDER"), uuid.String()),
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.Blob(http.StatusOK, mime, preview)
}
