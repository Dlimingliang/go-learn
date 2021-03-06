package gov1

import (
	"sync"
)

type URLStore struct {
	urls map[string]string
	mu   sync.Mutex
}

func NewURLStore() *URLStore {
	return &URLStore{urls: make(map[string]string)}
}

func (s *URLStore) Get(key string) string {

	s.mu.Lock()
	defer s.mu.Unlock()
	return s.urls[key]
}

func (s *URLStore) Set(key, url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, present := s.urls[key]; present {
		return false
	}
	s.urls[key] = url
	return true
}

func (s *URLStore) Count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.urls)
}

func (s *URLStore) Put(url string) string {
	for {
		key := genKey(s.Count())
		if ok := s.Set(key, url); ok {
			return key
		}
	}
	return ""
}
