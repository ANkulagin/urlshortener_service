package app

import (
	"github.com/ANkulagin/urlshortener_service/internal/app/routes"
	"net/http"
)

// RunServer запускает сервер
func RunServer() {
	r := routes.NewRouter()
	http.ListenAndServe(":8080", r)
}
