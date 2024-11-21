package handlers

import (
	"net/http"

	"github.com/AnonymFromInternet/EffectiveMobile/internal/repository"
)

func GETLibData(storage repository.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GET LIB DATA ROUTE"))
	}
}

func GETSongText(storage repository.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GET SONG TEXT ROUTE"))
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

func POSTAddNewSong(storage repository.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("POST NEW SONG ROUTE"))
	}
}
