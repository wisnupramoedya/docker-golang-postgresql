package model

type Movie struct {
	Title       string `json:"title"`
	Overview    string `json:"overview"`
	ReleaseDate string `json:"releaseDate"`
	Genres      string `json:"genres"`
}
