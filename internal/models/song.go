package models

type Song struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
}
