package handlers

import (
	"github.com/ANkulagin/urlshortener_service/internal/app/shortener"
	"github.com/ANkulagin/urlshortener_service/internal/app/storage"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

var urlStore = storage.NewURLStorage()

// ShortenURL обрабатывает POST-запрос для сокращения URL
func ShortenURL(w http.ResponseWriter, r *http.Request) {
	log := logrus.WithFields(logrus.Fields{
		"method": r.Method,
		"path":   r.URL.Path,
	})

	log.Info("Processing ShortenURL request")

	// Получение данных из тела POST-запроса
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("Error reading request body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Ваша логика для сокращения URL
	originalURL := string(body)
	shortURL := shortener.GenerateShortURL()

	// Сохранение сопоставления сокращенного URL с оригинальным
	urlStore.Save(shortURL, originalURL)

	// Отправка сокращенного URL в ответе
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortURL))
}

// RedirectURL обрабатывает GET-запрос для редиректа на оригинальный URL
func RedirectURL(w http.ResponseWriter, r *http.Request) {
	log := logrus.WithFields(logrus.Fields{
		"method": r.Method,
		"path":   r.URL.Path,
	})

	log.Info("Processing RedirectURL request")

	// Извлечение параметра id из пути
	id := chi.URLParam(r, "id")

	// Ваша логика для получения оригинального URL по id
	originalURL := urlStore.Get(id)

	if originalURL == "" {
		log.Error("Original URL not found for id:", id)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// Редирект на оригинальный URL
	http.Redirect(w, r, originalURL, http.StatusTemporaryRedirect)
}
