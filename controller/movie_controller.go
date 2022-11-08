package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/wisnupramoedya/movie-app/repository"
	"net/http"
)

func GetAllMovies(c echo.Context) error {
	movies, _ := repository.GetAllMoviesDB()
	return c.JSON(http.StatusOK, movies)
}
