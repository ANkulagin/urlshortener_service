package routes

import (
	"github.com/ANkulagin/urlshortener_service/internal/app/handlers"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"net/http"
)

// NewRouter создает новый роутер с обработчиками
func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(LoggerMiddleware)
	r.Post("/", handlers.ShortenURL)
	r.Get("/{id}", handlers.RedirectURL)
	return r
}

// LoggerMiddleware возвращает middleware для логирования запросов
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{
			"method": r.Method,
			"path":   r.URL.Path,
		}).Info("Request received")
		next.ServeHTTP(w, r)
	})
}
