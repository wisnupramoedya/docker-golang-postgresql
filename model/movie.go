package model

import (
	"github.com/lib/pq"
	"time"
)

type Movie struct {
	Title       string         `json:"title"`
	Overview    string         `json:"overview"`
	ReleaseDate time.Time      `json:"releaseDate"`
	Genres      pq.StringArray `gorm:"type:text[]" json:"genres"`
}
