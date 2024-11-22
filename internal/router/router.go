package router

import (
	"log/slog"

	"github.com/AnonymFromInternet/EffectiveMobile/internal/handlers"
	"github.com/AnonymFromInternet/EffectiveMobile/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New(storage repository.Repository, externalApiUrl string, logger *slog.Logger) *chi.Mux {
	mux := chi.NewMux()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.URLFormat)
	mux.Use(middleware.RealIP)

	mux.Route("/songs", func(r chi.Router) {
		r.Get("/", handlers.GETAllSongs(storage, logger))
		r.Get("/{id}", handlers.GETSongText(storage, logger))

		r.Delete("/{id}", handlers.DELETESong(storage))

		r.Patch("/{id}", handlers.PATCHSong(storage))

		r.Post("/", handlers.POSTNewSong(storage, externalApiUrl))
	})

	return mux
}
