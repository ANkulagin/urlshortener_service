package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRunServer(t *testing.T) {
	// Создаем фейковый ResponseWriter
	w := httptest.NewRecorder()

	// Создаем фейковый Request (не важен его контент, так как мы проверяем только запуск сервера)
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	// Запускаем сервер (это блокирующий вызов, поэтому запускаем его в горутине)
	go RunServer()

	// Ждем, чтобы сервер успел запуститься
	time.Sleep(time.Millisecond * 100)

	// Вызываем функцию handleRequest напрямую с фейковыми w и r
	handleRequest(w, r)

	// Проверяем код ответа
	assert.Equal(t, http.StatusMethodNotAllowed, w.Code, "ожидается код ответа 405 Method Not Allowed")
}

func Test_handleRequest(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		expectedCode int
		expectedBody string
	}{
		{name: "PostRequest", method: http.MethodPost, expectedCode: http.StatusCreated, expectedBody: "http://localhost:8080/"},
		{name: "GetRequest", method: http.MethodGet, expectedCode: http.StatusTemporaryRedirect},
		{name: "InvalidRequest", method: http.MethodPut, expectedCode: http.StatusMethodNotAllowed},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем фейковый ResponseWriter
			w := httptest.NewRecorder()

			// Создаем фейковый Request с указанным методом
			r := httptest.NewRequest(tt.method, "/", nil)

			// Вызываем функцию для тестирования
			handleRequest(w, r)

			// Проверяем код ответа
			assert.Equal(t, tt.expectedCode, w.Code, "код ответа не совпадает с ожиданиями")

			// Проверяем тело ответа, если оно ожидается
			if tt.expectedBody != "" {
				assert.Equal(t, tt.expectedBody, w.Body.String(), "тело ответа не совпадает с ожидаемым")
			}
		})
	}
}
