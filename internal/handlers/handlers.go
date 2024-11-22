package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/AnonymFromInternet/EffectiveMobile/internal/filter"
	"github.com/AnonymFromInternet/EffectiveMobile/internal/handlers/helpers"
	"github.com/AnonymFromInternet/EffectiveMobile/internal/models"
	"github.com/AnonymFromInternet/EffectiveMobile/internal/repository"
	"github.com/go-chi/chi/v5"

	loggerPackage "github.com/AnonymFromInternet/EffectiveMobile/internal/logger"
)

type response struct {
	Status  int           `json:"status"`
	Error   bool          `json:"error"`
	Message string        `json:"message"`
	Songs   []models.Song `json:"songs"`
}

func getResponse(status int, e bool, songs []models.Song, message string) response {
	return response{
		Error:   e,
		Status:  status,
		Songs:   songs,
		Message: message,
	}
}

func GETAllSongs(storage repository.Repository, logger *slog.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		filterType := chi.URLParam(r, "filterType")
		filterValue := chi.URLParam(r, "filterValue")
		skip := chi.URLParam(r, "skip")
		skipAsInt, e := strconv.Atoi(skip)
		if e != nil {
			logger.Error("package handlers.GETAllSongs: cannot parse skip into int ", loggerPackage.WrapError(e))
			helpers.SendResponse(w, 500, getResponse(500, true, []models.Song{}, ""))
			return
		}

		rows := chi.URLParam(r, "rows")
		rowsAsInt, e := strconv.Atoi(rows)
		if e != nil {
			logger.Error("package handlers.GETAllSongs: cannot parse rows amount into int ", loggerPackage.WrapError(e))
			helpers.SendResponse(w, 500, getResponse(500, true, []models.Song{}, ""))
			return
		}

		f := struct {
			FType  string
			FValue string
			Skip   int
			Rows   int
		}{}

		if !filter.IsTypeValid(f) {
			helpers.SendResponse(w, 400, getResponse(400, true, []models.Song{}, fmt.Sprintf("Cannot use this filter type: %s. See swagger for available values", f.FType)))
			return
		}

		f.FType = filterType
		f.FValue = filterValue
		f.Skip = skipAsInt
		f.Rows = rowsAsInt

		songs, e := storage.GetSongs(f)
		if e != nil {
			logger.Error("package handlers.GETAllSongs: cannot get songs from database", loggerPackage.WrapError(e))
			helpers.SendResponse(w, 500, getResponse(500, true, []models.Song{}, ""))
			return
		}

		helpers.SendResponse(w, 200, getResponse(200, false, songs, ""))
	}
}

func GETSongText(storage repository.Repository, logger *slog.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		skip := chi.URLParam(r, "skip")
		skipAsInt, e := strconv.Atoi(skip)
		if e != nil {
			logger.Error("package handlers.GETSong: cannot parse skip into int ", loggerPackage.WrapError(e))
			helpers.SendResponse(w, 500, getResponse(500, true, []models.Song{}, ""))
			return
		}

		rows := chi.URLParam(r, "rows")
		rowsAsInt, e := strconv.Atoi(rows)
		if e != nil {
			logger.Error("package handlers.GETSong: cannot parse rows amount into int ", loggerPackage.WrapError(e))
			helpers.SendResponse(w, 500, getResponse(500, true, []models.Song{}, ""))
			return
		}

		id := chi.URLParam(r, "id")
		idAsInt, e := strconv.Atoi(id)
		if e != nil {
			logger.Error("package handlers.GETSong: cannot parse id into int ", loggerPackage.WrapError(e))
			helpers.SendResponse(w, 500, getResponse(500, true, []models.Song{}, ""))
			return
		}

		f := struct {
			FType  string
			FValue string
			Skip   int
			Rows   int
		}{}

		f.Skip = skipAsInt
		f.Rows = rowsAsInt
		songText, e := storage.GetSongText(idAsInt, f)
		if e != nil {
			logger.Error("package handlers.GETSong: cannot get song from database", loggerPackage.WrapError(e))
			helpers.SendResponse(w, 500, getResponse(500, true, []models.Song{}, ""))
			return
		}

		helpers.SendResponse(w, 200, getResponse(200, false, []models.Song{{Text: songText}}, ""))
	}
}

func DELETESong(storage repository.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("DELETE SONG ROUTE"))
	}
}

func PATCHSong(storage repository.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("PATCH SONG ROUTE"))
	}
}

func POSTNewSong(storage repository.Repository, externalApiUrl string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("POST NEW SONG ROUTE"))
	}
}

// Пока непонятно какие ручки делать
// По сваггеру надо GET "/info" - это скорее всего тогда отдельно стоящая ручка для получения всей информации
// У него есть параметры group , song

// Скорее всего нужно будет делать 2 таблицы: группа, песня

// /info?group="groupName"&song="songName" - этот сервис не надо поднимать. а просто сделать запрос к нему из post ручки и подобрать
// данные оттуда

// Отвечает в виде JSON
// {releaseDate: "16.07.2006", text: "Ooh baby, don't you know I suffer?\nOoh baby", link: "https://www.youtube.com/watch?v=Xsp3_a-PMTw"}

// Пагинацию возможно делать необязательным параметром. Если нет его, то выдавать всё. Если есть, то только запрошенное количество
// Вроде параметр называется skip
