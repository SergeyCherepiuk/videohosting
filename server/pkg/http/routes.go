package http

import (
	"os"

	"github.com/SergeyCherepiuk/videohosting/pkg/bucket"
	"github.com/SergeyCherepiuk/videohosting/pkg/http/handlers"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	S3Client *s3.Client
}

func (router Router) Build() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.BodyLimit("1G"))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderCacheControl},
	}))

	// Services
	s3BucketService := bucket.NewS3BucketService(router.S3Client, os.Getenv("S3_BUCKET_NAME"))

	// Handlers
	videoHandler := handlers.NewVideoHandler(s3BucketService)
	
	// Routes
	e.POST("/upload", videoHandler.Upload)

	return e
}
