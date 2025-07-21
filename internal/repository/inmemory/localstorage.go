package inmemory

import (
	"library_api/internal/domain/errors"
	"library_api/internal/domain/models"

	"github.com/google/uuid"
)

type LocalStorage struct {
	books map[string]models.Book
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		books: make(map[string]models.Book), // Инициализация хранилища книг
	}
}

func (ls *LocalStorage) GetBooksList() ([]models.Book, error) {
	var bookList []models.Book

	if len(ls.books) == 0 {
		return nil, errors.ErrBookListEmpty
	}

	// Преобразование карты в срез
	for _, book := range ls.books {
		bookList = append(bookList, book)
	}

	return bookList, nil
}

func (ls *LocalStorage) SaveBook(book models.Book) {

	for key, b := range ls.books {
		if b.Author == book.Author && b.Label == book.Label {
			// Обновление количества книг
			// Книга уже существует, просто увеличиваем количество
			mBook := ls.books[key]
			mBook.Count++
			ls.books[key] = mBook
			return
		}
	}
	bookID := uuid.New().String() // Генерация уникального ID для книги
	book.ID = bookID
	book.Count = 1          // Устанавливаем количество книг в 1 при создании новой книги
	ls.books[bookID] = book // Сохранение книги в хранилище
}
