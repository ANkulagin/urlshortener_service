package app

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_handleShortenURL(t *testing.T) {
	// Создаем фейковый ResponseWriter
	w := httptest.NewRecorder()

	// Создаем фейковый Request с телом запроса
	r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("https://example.com/long-url"))

	// Вызываем функцию для тестирования
	handleShortenURL(w, r)

	// Проверяем код состояния ответа
	assert.Equal(t, http.StatusCreated, w.Code, "код ответа не совпадает с ожиданиями")

	// Проверяем тип содержимого ответа
	assert.Equal(t, "text/plain", w.Header().Get("Content-Type"), "Content-Type не совпадает с ожиданиями")

	// Проверяем тело ответа
	expectedBody := "http://localhost:8080/"
	assert.Equal(t, expectedBody, w.Body.String(), "Тело ответа не совпадает с ожидаемым")
}

func Test_handleRedirect(t *testing.T) {
	// Создаем фейковый ResponseWriter
	w := httptest.NewRecorder()

	// Создаем фейковый Request с путем /EwHXdJfB
	r := httptest.NewRequest(http.MethodGet, "/EwHXdJfB", nil)

	// Вызываем функцию для тестирования
	handleRedirect(w, r)

	// Проверяем код состояния ответа
	assert.Equal(t, http.StatusTemporaryRedirect, w.Code, "код ответа не совпадает с ожиданиями")

	// Проверяем заголовок Location
	expectedLocation := "https://example.com/long-url"
	assert.Equal(t, expectedLocation, w.Header().Get("Location"), "Location не совпадает с ожиданиями")
}
