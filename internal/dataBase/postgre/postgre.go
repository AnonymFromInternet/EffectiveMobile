package postgre

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"

	"github.com/AnonymFromInternet/EffectiveMobile/internal/filter"
	loggerPackage "github.com/AnonymFromInternet/EffectiveMobile/internal/logger"
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
		logger.Error("package postgre.MustCreate: cannot create dbConnection", loggerPackage.WrapError(e))
		log.Fatal("package postgre.MustCreate: cannot create dbConnection")
	}

	migrations.Up(dbConnection, upMigrationPath)

	return &Storage{DB: dbConnection, Logger: logger}
}

func (s *Storage) GetSongsFiltered(f filter.Filter) ([]models.Song, error) {
	var filterStmt string
	if !filter.IsTypeEmpty(f) {
		switch f.FType {
		case "name":
			filterStmt += ` WHERE s.name=$3`
		case "group":
			filterStmt += ` WHERE a.group=$3`
		case "releaseDate":
			filterStmt += ` WHERE s.releaseDate=$3`
		case "text":
			filterStmt += ` WHERE s.text=$3`
		case "link":
			filterStmt += ` WHERE s.link=$3`
		}

	}

	baseStmt := `SELECT s.id, s.name, s.release_date, a.name as group, s.song_text, s.link FROM song s
				JOIN artist a
				ON s.artist_id = a.id
		`

	limitStmt := ` LIMIT $1 OFFSET $2`

	finalStmt := baseStmt + filterStmt + limitStmt
	stmt, e := s.DB.Prepare(finalStmt)
	if e != nil {
		return nil, e
	}

	rows, e := stmt.Query(f.Rows, f.Skip, f.FValue)
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

func (s *Storage) GetSongs() ([]models.Song, error) {
	rows, e := s.DB.Query(`SELECT * FROM song s
				JOIN artist a
				ON s.artist_id = a.id`)
	fmt.Println("rows :", rows)
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

func (s *Storage) GetSongText(id int) (string, error) {
	stmtText := `SELECT song_text FROM song
				WHERE id = $1
			`
	stmt, e := s.DB.Prepare(stmtText)
	if e != nil {
		return "", e
	}

	res := stmt.QueryRow(id)

	var songText string
	e = res.Scan(&songText)
	if e != nil {
		return "", e
	}

	return songText, nil
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

func (s *Storage) AddSong(song models.Song, groupId int) error {
	stmtText := `INSERT INTO song(name, release_date, artist_id, song_text, link) values($1, $2, $3, $4, $5)`

	stmt, e := s.DB.Prepare(stmtText)
	if e != nil {
		return e
	}

	_, e = stmt.Exec(song.Name, song.ReleaseDate, groupId, song.Text, song.Link)
	if e != nil {
		return e
	}

	return nil
}

func (s *Storage) GetGroup(name string) (int, error) {
	basicStmt := `SELECT id from artist WHERE name = $1`
	stmt, _ := s.DB.Prepare(basicStmt)

	groupId := -1
	res := stmt.QueryRow(name)
	e := res.Scan(&groupId)

	return groupId, e
}

func (s *Storage) AddGroup(name string) (int, error) {
	basicStmt := `INSERT INTO artist(name) values($1) RETURNING id`
	stmt, _ := s.DB.Prepare(basicStmt)
	var id int
	newGroupId := stmt.QueryRow(name)

	e := newGroupId.Scan(&id)
	if e != nil {
		return -1, e
	}

	return id, nil
}
