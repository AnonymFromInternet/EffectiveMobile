package postgre

import (
	"database/sql"
	"log/slog"

	"github.com/AnonymFromInternet/EffectiveMobile/internal/models"
)

type Storage struct {
	DB     *sql.DB
	Logger *slog.Logger
}

func MustCreate(dataSourceName string, logger *slog.Logger) *Storage {
	// Тут запустить миграцию по созданию базы
	// Нужно только одну? Какие вообще миграции нужно выполнять?

	return &Storage{Logger: logger}
}

func (s *Storage) GetData() (models.Data, error) {
	return models.Data{}, nil
}

func (s *Storage) GetSongText(id int) (string, error) {
	return "", nil
}

func (s *Storage) DeleteSong(id int) error {
	return nil
}

func (s *Storage) ChangeSong(id int, changedSong models.Song) error {
	return nil
}

func (s *Storage) AddSong(song models.Song) error {
	return nil
}
