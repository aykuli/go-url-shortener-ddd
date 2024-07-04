package repository

import (
	"fmt"

	"github.com/vokinneberg/go-url-shortener-ddd/domain"
)

type Store interface {
	Add(url *domain.URL) error
	Read(id string) (*domain.URL, error)
}

type MemStore struct {
	items map[string]string
}

func NewMemStore() *MemStore {
	return &MemStore{items: make(map[string]string)}
}

func (m *MemStore) Add(url *domain.URL) error {
	m.items[url.ID] = url.Original

	return nil
}

func (m *MemStore) Read(id string) (*domain.URL, error) {
	if _, ok := m.items[id]; ok {
		return nil, fmt.Errorf("can't find requested id %s", id)
	}

	return &domain.URL{Original: m.items[id], ID: id}, nil
}
