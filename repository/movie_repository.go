package repository

import (
	"github.com/wisnupramoedya/movie-app/config"
	"github.com/wisnupramoedya/movie-app/model"
)

func GetAllMoviesDB() ([]model.Movie, error) {
	db := config.GetDBInstance()
	var movies []model.Movie

	if err := db.Find(&movies).Error; err != nil {
		return nil, err
	}

	return movies, nil
}
