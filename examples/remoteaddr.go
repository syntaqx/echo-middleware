package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/syntaqx/echo-middleware/remoteaddr"
)

func main() {
	e := echo.New()

	e.Use(remoteaddr.New().Handler)

	e.Get("/", func(c *echo.Context) error {
		return c.HTML(http.StatusOK, c.Request().RemoteAddr)
	})

	e.Run(":8080")
}
