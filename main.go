package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
	})

	e.Start(":3000")
}
