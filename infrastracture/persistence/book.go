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
	if err := r.Conn.Debug().Create(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (r *bookRepository) FetchBooks() ([]bookdm.Book, error) {
	var books []bookdm.Book

	if err := r.Conn.Debug().Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *bookRepository) SearchBooks(title string, author string) ([]bookdm.Book, error) {
	var books []bookdm.Book
	if err := r.Conn.Debug().Where("title like ?", "%"+title+"%").Or("author like ?", "%"+author+"%").Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}
