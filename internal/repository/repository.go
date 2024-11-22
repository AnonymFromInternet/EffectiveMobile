package repository

import (
	"github.com/AnonymFromInternet/EffectiveMobile/internal/filter"
	"github.com/AnonymFromInternet/EffectiveMobile/internal/models"
)

type Repository interface {
	GetSongs(filter filter.Filter) ([]models.Song, error)
	GetSongText(id int, f filter.Filter) (string, error)
	DeleteSong(id int) error
	ChangeSong(id int, changedSong models.Song) error
	AddSong(song models.Song) error
}
