package repository

import (
	"github.com/wisnupramoedya/movie-app/config"
	"github.com/wisnupramoedya/movie-app/model"
)

func GetAllMoviesDB() ([]model.Movie, error) {
	db := config.GetDBInstance()
	var movies []model.Movie

	result := db.Find(&movies)
	return movies, result.Error
}

func GetMovieDBWithId(id int) (model.Movie, error) {
	db := config.GetDBInstance()
	var movie model.Movie

	result := db.Where("index = ?", id).First(&movie)
	return movie, result.Error
}
