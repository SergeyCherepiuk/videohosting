package http

import (
	"github.com/SergeyCherepiuk/videohosting/pkg/bucket"
	"github.com/SergeyCherepiuk/videohosting/pkg/http/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
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
	videoHandler := handlers.NewVideoHandler(router.Bucket)

	// Routes
	e.POST("/upload", videoHandler.Upload)

	return e
}
