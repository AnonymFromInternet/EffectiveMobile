package repository

import "github.com/AnonymFromInternet/EffectiveMobile/internal/models"

type Repository interface {
	GetData() error
	GetSongText(id int) error
	DeleteSong(id int) error
	ChangeSong(id int, changedSong models.Song) error
	AddSong(song models.Song) error
}
