package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wisnupramoedya/movie-app/config"
	"github.com/wisnupramoedya/movie-app/controller"
)

func main() {
	e := echo.New()
	config.OpenConnection()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", controller.GetAllMovies)

	e.Logger.Fatal(e.Start(":9000"))
}
