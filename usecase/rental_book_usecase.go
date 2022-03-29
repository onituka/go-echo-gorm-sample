package usecase

import (
	"time"

	"github.com/onituka/go-echo-gorm-sample/domain/rentalbookdm"
)

type RentalBookUsecase interface {
	FetchRentalBook(rentalID int) (*rentalbookdm.RentalBook, error)
	CreateRentalBook(rentalBook *rentalbookdm.RentalBook) (*rentalbookdm.RentalBook, error)
}

type rentalBookUsecase struct {
	rentalBookRepository rentalbookdm.RentalBookRepository
}

func NewRentalBookUsecase(rentalBookRepository rentalbookdm.RentalBookRepository) *rentalBookUsecase {
	return &rentalBookUsecase{rentalBookRepository: rentalBookRepository}
}

func (u *rentalBookUsecase) FetchRentalBook(rentalID int) (*rentalbookdm.RentalBook, error) {
	rentalBook, err := u.rentalBookRepository.FetchRentalBook(rentalID)
	if err != nil {
		return nil, err
	}

	return rentalBook, nil
}

func (u *rentalBookUsecase) CreateRentalBook(rentalBook *rentalbookdm.RentalBook) (*rentalbookdm.RentalBook, error) {

	rentalDay := time.Now().UTC()

	deadLine := time.Now().UTC().AddDate(0, 0, 7)

	rentalBook.LoanDate = rentalDay
	rentalBook.ReturnDeadline = deadLine

	rentalDate := rentalbookdm.NewAddRentalBook(rentalBook.ID, rentalBook.BookID, rentalBook.UserID, rentalBook.LoanDate, rentalBook.ReturnDate, rentalBook.ReturnDeadline)

	createdRentalDate, err := u.rentalBookRepository.CreateRentalBook(rentalDate)
	if err != nil {
		return nil, err
	}

	return createdRentalDate, err
}
