package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/AnonymFromInternet/EffectiveMobile/internal/filter"
	"github.com/AnonymFromInternet/EffectiveMobile/internal/handlers/helpers"
	"github.com/AnonymFromInternet/EffectiveMobile/internal/models"
	"github.com/AnonymFromInternet/EffectiveMobile/internal/repository"
	"github.com/go-chi/chi/v5"

	loggerPackage "github.com/AnonymFromInternet/EffectiveMobile/internal/logger"
)

func GETAllSongs(storage repository.Repository, logger *slog.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		filterType := r.URL.Query().Get("filterType")
		filterValue := r.URL.Query().Get("filterValue")
		skip := r.URL.Query().Get("skip")
		rows := r.URL.Query().Get("rows")
		skipAsInt, e := strconv.Atoi(skip)
		if e != nil {
			logger.Error("package handlers.GETAllSongs: cannot parse skip into int ", loggerPackage.WrapError(e))
			helpers.SendResponse(w, 500, helpers.GetResponse(500, true, []models.Song{}, ""))
			return
		}

		rowsAsInt, e := strconv.Atoi(rows)
		if e != nil {
			logger.Error("package handlers.GETAllSongs: cannot parse rows amount into int ", loggerPackage.WrapError(e))
			helpers.SendResponse(w, 500, helpers.GetResponse(500, true, []models.Song{}, ""))
			return
		}

		f := struct {
			FType  string
			FValue string
			Skip   int
			Rows   int
		}{}
		f.FType = filterType
		f.FValue = filterValue
		f.Skip = skipAsInt
		f.Rows = rowsAsInt

		if filter.IsClear(f) {
			songs, e := storage.GetSongs()
			if e != nil {
				logger.Error("package handlers.GETAllSongs: cannot get songs from database", loggerPackage.WrapError(e))
				helpers.SendResponse(w, 500, helpers.GetResponse(500, true, []models.Song{}, ""))
				return
			}

			helpers.SendResponse(w, 200, helpers.GetResponse(200, false, songs, ""))
		} else {
			if !filter.IsTypeValid(f) {
				logger.Debug("package handlers.GETAllSongs: invalid filter type from client", loggerPackage.WrapDebug(f.FType))
				helpers.SendResponse(w, 400, helpers.GetResponse(400, true, []models.Song{}, fmt.Sprintf("Cannot use this filter type: %s. See swagger for available values", f.FType)))
				return
			}

			songs, e := storage.GetSongsFiltered(f)
			if e != nil {
				logger.Error("package handlers.GETAllSongs: cannot get songs from database", loggerPackage.WrapError(e))
				helpers.SendResponse(w, 500, helpers.GetResponse(500, true, []models.Song{}, ""))
				return
			}

			helpers.SendResponse(w, 200, helpers.GetResponse(200, false, songs, ""))
		}
	}
}

func GETSongText(storage repository.Repository, logger *slog.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		skip := r.URL.Query().Get("skip")
		skipAsInt, e := strconv.Atoi(skip)
		if e != nil {
			logger.Error("package handlers.GETSong: cannot parse skip into int ", loggerPackage.WrapError(e))
			helpers.SendResponse(w, 500, helpers.GetResponse(500, true, []models.Song{}, ""))
			return
		}

		id := chi.URLParam(r, "id")
		idAsInt, e := strconv.Atoi(id)
		if e != nil {
			logger.Error("package handlers.GETSong: cannot parse id into int ", loggerPackage.WrapError(e))
			helpers.SendResponse(w, 500, helpers.GetResponse(500, true, []models.Song{}, ""))
			return
		}

		songText, e := storage.GetSongText(idAsInt)
		if e != nil {
			logger.Error("package handlers.GETSong: cannot get song from database", loggerPackage.WrapError(e))
			helpers.SendResponse(w, 500, helpers.GetResponse(500, true, []models.Song{}, ""))
			return
		}

		songText = strings.ReplaceAll(songText, "\\n", "\n")
		slicedText := strings.Split(songText, "\n")
		if skipAsInt > len(slicedText)-1 {
			helpers.SendResponse(w, 400, helpers.GetResponse(400, false, []models.Song{{}}, "false skip value for this song"))
			logger.Debug("package handlers.GETSong: false skip value for song:", loggerPackage.WrapDebug("song id is "+id))
			return
		}

		res := slicedText[skipAsInt:]

		songText = strings.Join(res, "\n")

		helpers.SendResponse(w, 200, helpers.GetResponse(200, false, []models.Song{{Text: songText}}, ""))
	}
}

func DELETESong(storage repository.Repository, logger *slog.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idAsInt, e := strconv.Atoi(id)
		if e != nil {
			logger.Error("package handlers.GETSong: cannot parse id into int ", loggerPackage.WrapError(e))
			helpers.SendResponse(w, 500, helpers.GetResponse(500, true, []models.Song{}, ""))
			return
		}

		e = storage.DeleteSong(idAsInt)
		if e != nil {
			logger.Error("package handlers.GETSong: cannot delete song", loggerPackage.WrapError(e))
			helpers.SendResponse(w, 500, helpers.GetResponse(500, true, []models.Song{}, ""))
			return
		}

		helpers.SendResponse(w, 200, nil)
	}
}

func PATCHSong(storage repository.Repository, logger *slog.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idAsInt, e := strconv.Atoi(id)
		if e != nil {
			logger.Error("package handlers.PATCHSong: cannot convert id into int", loggerPackage.WrapError(e))
			helpers.SendResponse(w, 500, helpers.GetResponse(500, true, []models.Song{}, ""))
			return
		}

		// Получить пейлоад
		var payload models.Song
		e = json.NewDecoder(r.Body).Decode(&payload)
		defer r.Body.Close()
		if e != nil {
			logger.Error("package handlers.PATCHSong: cannot deserialize payload", loggerPackage.WrapError(e))
			helpers.SendResponse(w, 500, helpers.GetResponse(500, true, []models.Song{}, ""))
			return
		}

		// Вызвать метод базы данных
		e = storage.ChangeSong(idAsInt, payload)
		if e != nil {
			logger.Error("package handlers.PATCHSong: cannot change song in database", loggerPackage.WrapError(e))
			helpers.SendResponse(w, 500, helpers.GetResponse(500, true, []models.Song{}, ""))
			return
		}

		helpers.SendResponse(w, 204, helpers.GetResponse(204, true, []models.Song{}, ""))
	}
}

func POSTNewSong(storage repository.Repository, externalApiUrl string, logger *slog.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload helpers.Payload
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		if e := decoder.Decode(&payload); e != nil {
			logger.Error("package handlers.POSTNewSong: cannot decode payload", loggerPackage.WrapError(e))
			helpers.SendResponse(w, 500, helpers.GetResponse(500, true, []models.Song{}, ""))
			return
		}

		resp, e := http.Get(externalApiUrl + "?group=" + payload.Group + "&song=" + payload.Song)
		if e != nil {
			logger.Error("package handlers.POSTNewSong: cannot call external api", loggerPackage.WrapError(e))
			helpers.SendResponse(w, 500, helpers.GetResponse(500, true, []models.Song{}, ""))
			return
		}

		decoder = json.NewDecoder(resp.Body)
		defer resp.Body.Close()

		var song models.Song
		e = decoder.Decode(&song)
		if e != nil {
			logger.Error("package handlers.POSTNewSong: cannot deserialize external resp into song", loggerPackage.WrapError(e))
			helpers.SendResponse(w, 500, helpers.GetResponse(500, true, []models.Song{}, ""))
			return
		}

		song.Group = payload.Group
		song.Name = payload.Song
		song.ReleaseDate, e = helpers.ConvertDate(song.ReleaseDate)
		if e != nil {
			logger.Error("package handlers.POSTNewSong: cannot convert date", loggerPackage.WrapError(e))
			helpers.SendResponse(w, 500, helpers.GetResponse(500, true, []models.Song{}, ""))
			return
		}

		groupId, e := storage.GetGroup(song.Name)
		isNoGroup := e == nil && groupId == -1
		if isNoGroup {
			groupId, e = storage.AddGroup(song.Group)
			if e != nil {
				logger.Error("package handlers.POSTNewSong: cannot add new group", loggerPackage.WrapError(e))
				helpers.SendResponse(w, 500, helpers.GetResponse(500, true, []models.Song{}, ""))
				return
			}
		}

		e = storage.AddSong(song, groupId)
		if e != nil {
			logger.Error("package handlers.POSTNewSong: cannot save new song into database", loggerPackage.WrapError(e))
			helpers.SendResponse(w, 500, helpers.GetResponse(500, true, []models.Song{}, ""))
			return
		}

		helpers.SendResponse(w, 204, helpers.GetResponse(204, true, []models.Song{}, "new song added"))
	}
}
