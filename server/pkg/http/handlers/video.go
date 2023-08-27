package handlers

import (
	"context"
	"io"
	"net/http"

	"github.com/SergeyCherepiuk/videohosting/pkg/bucket"
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

	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if err := handler.bucket.Upload(context.Background(), fileHeader.Filename, content); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
