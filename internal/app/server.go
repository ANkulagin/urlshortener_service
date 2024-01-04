package app

import (
	"fmt"
	"github.com/ANkulagin/urlshortener_service/internal/app/routes"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

// RunServer запускает сервер
func RunServer(port string) {
	// Настройка Logrus
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)

	// Формирование адреса для прослушивания
	addr := fmt.Sprintf(":%s", port)

	r := routes.NewRouter()
	fmt.Println("Running server on", addr)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		return
	}
}
