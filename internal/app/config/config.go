package config

import "github.com/spf13/cobra"

// Config представляет структуру конфигурации сервиса
type Config struct {
	HTTPAddr     string // Адрес HTTP-сервера
	BaseShortURL string // Базовый адрес для сокращенных URL
}

var config Config // Глобальная переменная для хранения текущей конфигурации

// InitializeConfig инициализирует конфигурацию из аргументов командной строки
func InitializeConfig(cmd *cobra.Command) {
	// Извлекаем значения флагов из команды Cobra и присваиваем их соответствующим полям конфигурации
	config.HTTPAddr, _ = cmd.Flags().GetString("http-addr")
	config.BaseShortURL, _ = cmd.Flags().GetString("base-short-url")
}

// GetConfig возвращает текущую конфигурацию
func GetConfig() *Config {
	return &config // Возвращаем указатель на текущую конфигурацию
}
