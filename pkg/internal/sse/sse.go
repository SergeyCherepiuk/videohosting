package sse

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Event struct {
	ID        int
	EventName string
	Data      string
}

func (e Event) Format() []byte {
	return []byte(fmt.Sprintf("id: %d\nevent: %s\ndata: %s\n\n", e.ID, e.EventName, e.Data))
}

func (e Event) Send(response *echo.Response) error {
	_, err := response.Write(e.Format())
	response.Flush()
	return err
}
