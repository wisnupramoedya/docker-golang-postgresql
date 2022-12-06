package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/wisnupramoedya/movie-app/model"
	"github.com/wisnupramoedya/movie-app/repository"
	"net/http"
	"strconv"
)

func GetAllMovies(c echo.Context) error {
	response := new(model.Response)
	movies, err := repository.GetAllMoviesDB()
	if err != nil {
		response.ErrorCode = 404
		response.Message = "Gagal melihat data movies"
		return c.JSON(http.StatusNotFound, response)
	}

	response.ErrorCode = 0
	response.Message = "Sukses melihat data movies"
	response.Data = movies
	return c.JSON(http.StatusOK, response)
}

func GetMovieWithId(c echo.Context) error {
	response := new(model.Response)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ErrorCode = 404
		response.Message = "Request id bukanlah angka"
		return c.JSON(http.StatusBadRequest, response)
	}
	movie, _ := repository.GetMovieDBWithId(id)
	if err != nil {
		response.ErrorCode = 404
		response.Message = "Gagal melihat data movie"
		return c.JSON(http.StatusNotFound, response)
	}

	response.ErrorCode = 0
	response.Message = "Sukses melihat data movie"
	response.Data = movie
	return c.JSON(http.StatusOK, response)
}
