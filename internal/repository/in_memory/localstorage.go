package inmemory

import "library_api/internal/domain/models"

type LocalStorage struct {
	books map[string]models.Book
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		books: make(map[string]models.Book),
	}
}
