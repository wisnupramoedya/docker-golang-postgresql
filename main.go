package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/", func(ctx echo.Context) error {
		data := "Program web-app from the /index"
		return ctx.String(http.StatusOK, data)
	})

	e.Logger.Fatal(e.Start(":9000"))
}
