package http

import (
	"github.com/SergeyCherepiuk/videohosting/domain"
	"github.com/SergeyCherepiuk/videohosting/pkg/bucket"
	"github.com/SergeyCherepiuk/videohosting/pkg/http/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	VideoService domain.VideoService
	Bucket bucket.BucketService
}

func (router Router) Build() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.BodyLimit("1G"))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderCacheControl},
	}))

	// Handlers
	videoHandler := handlers.NewVideoHandler(router.VideoService, router.Bucket)
	previewHandler := handlers.NewPreviewHandler(router.Bucket)

	// Routes
	videos := e.Group("videos")
	videos.GET("/:uuid", videoHandler.GetByUUID)
	videos.POST("/upload", videoHandler.Upload)

	previews := e.Group("previews")
	previews.GET("/:uuid", previewHandler.GetByUUID)

	return e
}
