package router

import (
	"github.com/AnonymFromInternet/EffectiveMobile/internal/handlers"
	"github.com/AnonymFromInternet/EffectiveMobile/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New(storage repository.Repository) *chi.Mux {
	mux := chi.NewMux()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.URLFormat)

	// TODO: А libData сначала понять что это вообще такое, и потом уже понять либо сюда же, либо отдельно

	// TODO: установить обработчик на случай, если нет пути
	mux.Get("/libData", handlers.GETLibData(storage))

	mux.Route("/song", func(r chi.Router) {
		r.Get("/text", handlers.GETSongText(storage))

		r.Delete("/delete", handlers.DELETESong(storage))

		r.Patch("/change", handlers.PATCHSong(storage))

		r.Post("/add", handlers.POSTAddNewSong(storage))
	})

	return mux
}
