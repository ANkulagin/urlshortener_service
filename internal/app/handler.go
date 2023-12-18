package app

import (
	"fmt"
	"io"
	"net/http"
)

// URLMap - карта для хранения соответствия сокращенного URL и оригинального URL
var URLMap = make(map[string]string)

func handleShortenURL(w http.ResponseWriter, r *http.Request) {
	// Чтение тела запроса (URL) из тела запроса
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// В случае ошибки чтения возвращаем ошибку
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// Получение оригинального URL из тела запроса
	originalURL := string(body)
	// Генерация сокращенного URL
	shortURL := generateShortURL()

	// Сохранение соответствия сокращенного и оригинального URL
	URLMap[shortURL] = originalURL

	// Формирование ответа с сокращенным URL
	response := fmt.Sprintf("http://localhost:8080/%s", shortURL)

	// Установка статуса Created (201) и типа содержимого text/plain
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	// Отправка ответа
	w.Write([]byte(response))
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	// Извлечение идентификатора из URL (удаление ведущего слэша)
	id := r.URL.Path[1:]

	// Поиск оригинального URL в карте
	originalURL, exists := URLMap[id]
	if !exists {
		// В случае отсутствия соответствия возвращаем ошибку Not Found (404)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// Установка заголовка Location с оригинальным URL
	w.Header().Set("Location", originalURL)
	// Установка статуса Temporary Redirect (307)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
