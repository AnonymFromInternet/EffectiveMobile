package models

type Song struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	ReleaseDate string `json:"releaseDate"`
	Group       string `json:"group"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}
