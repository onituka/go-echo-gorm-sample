package usecase

import (
	"github.com/onituka/go-echo-gorm-sample/domain/bookdm"
)

type BookUsecase interface {
	FetchBook(bookID int) (*bookdm.Book, error)
}

type bookUsecase struct {
	bookRepository bookdm.BookRepository
}

func NewBookUsecase(bookRepository bookdm.BookRepository) *bookUsecase {
	return &bookUsecase{bookRepository: bookRepository}
}

func (u *bookUsecase) FetchBook(bookID int) (*bookdm.Book, error) {
	book, err := u.bookRepository.FetchBook(bookID)
	if err != nil {
		return nil, err
	}
	return book, nil
}
