package http

import (
	"fmt"
	"io"
	"log"
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

		fileHeader, err := c.FormFile("file")
		if err != nil {
			log.Println(err.Error())
			return err
		}

		file, err := fileHeader.Open()
		if err != nil {
			log.Println(err.Error())
			return err
		}

		content, err := io.ReadAll(file)
		if err != nil {
			log.Println(err.Error())
			return err
		}

		log.Printf("file length: %d", len(content))

		for i := 0; i < 5; i++ {
			fmt.Fprintf(
				c.Response().Writer,
				"id: %d\nevent: %s\ndata: %v",
				i, "test", i,
			)
			c.Response().Flush()
			time.Sleep(time.Second)
		}

		return c.NoContent(http.StatusOK)
	})

	return e
}
