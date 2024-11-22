package postgre

import (
	"database/sql"
	"log"
	"log/slog"

	"github.com/AnonymFromInternet/EffectiveMobile/internal/filter"
	"github.com/AnonymFromInternet/EffectiveMobile/internal/migrations"
	"github.com/AnonymFromInternet/EffectiveMobile/internal/models"
	_ "github.com/lib/pq"
)

type Storage struct {
	DB     *sql.DB
	Logger *slog.Logger
}

func MustCreate(dataSourceName string, logger *slog.Logger, upMigrationPath string) *Storage {
	dbConnection, e := sql.Open("postgres", dataSourceName)
	if e != nil {
		logger.Error("package postgre.MustCreate: cannot create dbConnection")
		log.Fatal("package postgre.MustCreate: cannot create dbConnection")
	}

	migrations.Up(dbConnection, upMigrationPath)

	return &Storage{DB: dbConnection, Logger: logger}
}

func (s *Storage) GetSongs(f filter.Filter) ([]models.Song, error) {
	var filterStmt string
	if !filter.IsEmpty(f) {
		if f.FType == "group" {
			filterStmt = ` WHERE a.name=$4`
		} else {
			filterStmt += ` WHERE s.$3=$4`
		}
	}

	baseStmt := `SELECT s.id, s.name, s.release_date, a.name as group, s.song_text, s.link FROM songs s
				JOIN artist a
				ON s.artist_id = a.id
				LIMIT $1 OFFSET $2
		`

	finalStmt := baseStmt + filterStmt
	stmt, e := s.DB.Prepare(finalStmt)
	if e != nil {
		return nil, e
	}

	rows, e := stmt.Query(f.Rows, f.Skip, f.FType, f.FValue)
	defer rows.Close()
	if e != nil {
		return nil, e
	}

	var songs []models.Song
	for rows.Next() {
		var song models.Song
		rows.Scan(&song.Id, &song.Name, &song.ReleaseDate, &song.Group, &song.Text, &song.Link)
		songs = append(songs, song)
	}

	return songs, nil
}

func (s *Storage) GetSongText(id int, f filter.Filter) (string, error) {
	// Пагинация по куплетам
	// Брать также rows и skip
	return "", nil
}

func (s *Storage) DeleteSong(id int) error {
	stmt, e := s.DB.Prepare(`DELETE FROM song WHERE id=$1`)
	if e != nil {
		return e
	}

	_, e = stmt.Exec(id)
	if e != nil {
		return e
	}
	return nil
}

func (s *Storage) ChangeSong(id int, changedSong models.Song) error {
	return nil
}

func (s *Storage) AddSong(song models.Song) error {
	return nil
}
