package app

import (
	"github.com/ANkulagin/urlshortener_service/internal/app/routes"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

// RunServer запускает сервер
func RunServer() {
	// Настройка Logrus
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)

	r := routes.NewRouter()
	http.ListenAndServe(":8080", r)
}
