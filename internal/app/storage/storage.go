package storage

import (
	"sync"
)

// URLStorage представляет собой хранилище URL
type URLStorage struct {
	mu          sync.RWMutex
	urlMappings map[string]string
}

// NewURLStorage создает новое хранилище URL
func NewURLStorage() *URLStorage {
	return &URLStorage{
		urlMappings: make(map[string]string),
	}
}

// Save сохраняет сопоставление сокращенного URL с оригинальным
func (s *URLStorage) Save(shortURL, originalURL string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.urlMappings[shortURL] = originalURL
}

// Get возвращает оригинальный URL по сокращенному
func (s *URLStorage) Get(shortURL string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urlMappings[shortURL]
}
