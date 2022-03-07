package persistence

import (
	"github.com/jinzhu/gorm"

	"github.com/onituka/go-echo-gorm-sample/domain/bookdm"
)

type bookRepository struct {
	Conn *gorm.DB
}

func NewBookRepository(conn *gorm.DB) bookdm.BookRepository {
	return &bookRepository{Conn: conn}
}

func (r *bookRepository) FetchBook(bookID int) (*bookdm.Book, error) {
	book := &bookdm.Book{ID: bookID}

	if err := r.Conn.Debug().First(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (r *bookRepository) CreateBook(book *bookdm.Book) (*bookdm.Book, error) {
	if err := r.Conn.Create(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}
