package repository

import (
	"github.com/AnonymFromInternet/EffectiveMobile/internal/filter"
	"github.com/AnonymFromInternet/EffectiveMobile/internal/models"
)

type Repository interface {
	GetSongs() ([]models.Song, error)
	GetSongsFiltered(filter filter.Filter) ([]models.Song, error)
	GetSongText(id int) (string, error)
	DeleteSong(id int) error
	ChangeSong(id int, changedSong models.Song) error
	AddSong(song models.Song, groupId int) error
	GetGroup(name string) (int, error)
	AddGroup(name string) (int, error)
}
