package main

import (
	"fmt"
	"github.com/ANkulagin/urlshortener_service/internal/app"
	"github.com/ANkulagin/urlshortener_service/internal/app/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "url-shortener",
	Short: "URL Shortener Service",
	Run: func(cmd *cobra.Command, args []string) {
		// Инициализация конфигурации
		config.InitializeConfig(cmd)

		// Запуск сервера с передачей порта из конфигурации
		app.RunServer(config.GetConfig().HTTPAddr)
	},
}

func init() {
	// Настройка флагов для команды
	rootCmd.Flags().StringP("http-addr", "a", "8080", "HTTP server port") // Используем порт 8080 по умолчанию
	rootCmd.Flags().StringP("base-short-url", "b", "http://localhost:8080/", "Base short URL")
}

func main() {
	// Запуск выполнения корневой команды
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
