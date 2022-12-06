package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wisnupramoedya/movie-app/config"
	"github.com/wisnupramoedya/movie-app/controller"
	"net/http"
)

func main() {
	e := echo.New()
	config.OpenConnection()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/movies", controller.GetAllMovies)
	e.GET("/movies/:id", controller.GetMovieWithId)

	e.Logger.Fatal(e.Start(":9000"))
}
