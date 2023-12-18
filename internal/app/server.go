package app

import "net/http"

func RunServer() {
	// Настройка обработчика запросов и запуск сервера на порту 8080
	http.HandleFunc("/", nil)
	http.ListenAndServe(":8080", nil)
}
