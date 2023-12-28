package handlers

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShortenURL(t *testing.T) {
	// Создаем фейковый запрос
	body := []byte("https://www.example.com/original-url")
	req, err := http.NewRequest("POST", "/shorten", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Создаем фейковый ResponseWriter
	res := httptest.NewRecorder()

	// Вызываем вашу функцию обработки запроса
	ShortenURL(res, req)

	// Проверяем статус код ответа
	assert.Equal(t, http.StatusCreated, res.Code, "Ожидался статус код 201, но получен %v", res.Code)

}
