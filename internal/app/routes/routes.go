package routes

import (
	"github.com/ANkulagin/urlshortener_service/internal/app/handlers"
	"github.com/go-chi/chi"
)

// NewRouter создает новый роутер с обработчиками
func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", handlers.ShortenURL)
	r.Get("/{id}", handlers.RedirectURL)
	return r
}
