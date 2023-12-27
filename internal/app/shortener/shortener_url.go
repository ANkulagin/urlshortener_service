package shortener

import (
	"math/rand"
	"strconv"
)

func GenerateShortURL() string {
	// Ваша логика для генерации сокращенного URL
	return "http://localhost:8080/" + strconv.Itoa(rand.Intn(1000))
}
