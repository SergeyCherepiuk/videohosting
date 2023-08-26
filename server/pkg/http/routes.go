package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.BodyLimit("1G"))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderCacheControl},
	}))

	e.POST("/upload", func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "text/event-stream")
		c.Response().Header().Set("Cache-Control", "no-cache")
		c.Response().Header().Set("Connection", "keep-alive")

		event := "test"
		for i := 0; i < 5; i++ {
			fmt.Fprintf(
				c.Response().Writer,
				"id: %d\nevent: %s\ndata: %v",
				i, event, i,
			)
			c.Response().Flush()
			time.Sleep(time.Second)
		}

		return c.NoContent(http.StatusOK)
	})

	return e
}
