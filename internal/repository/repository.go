package repository

import "github.com/AnonymFromInternet/EffectiveMobile/internal/models"

type Repository interface {
	GetData() (models.Data, error)
	GetSongText(id int) (string, error)
	DeleteSong(id int) error
	ChangeSong(id int, changedSong models.Song) error
	AddSong(song models.Song) error
}
