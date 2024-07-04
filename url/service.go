package url

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/vokinneberg/go-url-shortener-ddd/domain"
	"github.com/vokinneberg/go-url-shortener-ddd/internal/repository"
)

type URLService struct {
	store repository.Store
}

func NewURLService(store repository.Store) *URLService {
	return &URLService{
		store: store,
	}
}

func (h *URLService) Shorten(original string) (*domain.URL, error) {
	hash := sha1.New()
	hash.Write([]byte(original))
	short := hex.EncodeToString(hash.Sum(nil))[:8]
	return domain.NewURL(short, original), nil
}

func (h *URLService) Find(id string) (*domain.URL, error) {
	url, err := h.store.Read(id)
	if err != nil {
		return nil, err
	}

	return url, nil
}
