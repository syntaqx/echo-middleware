package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/rs/cors"
	"github.com/syntaqx/echo-middleware/method"
)

func main() {
	e := echo.New()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	e.Use(c.Handler)
	e.Use(method.Override())

	e.Get("/hello", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello Get")
	})

	e.Put("/hello", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello Put")
	})

	e.Run(":8080")
}
