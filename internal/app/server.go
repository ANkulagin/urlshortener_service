package app

import "net/http"

func RunServer() {
	// Настройка обработчика запросов и запуск сервера на порту 8080
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Обработка запроса POST для сокращения URL
		handleShortenURL(w, r)
	case http.MethodGet:
		// Обработка запроса GET для перенаправления на оригинальный URL
		handleRedirect(w, r)
	default:
		// В случае недопустимого метода возвращаем ошибку
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
