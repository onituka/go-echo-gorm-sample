package usecase

import (
	"time"

	"github.com/onituka/go-echo-gorm-sample/domain/rentalbookdm"
)

type RentalBookUsecase interface {
	FetchRentalBook(rentalID int) (*rentalbookdm.RentalBook, error)
	FetchRentalBooks() ([]rentalbookdm.RentalBook, error)
	CreateRentalBook(rentalBook *rentalbookdm.RentalBook) (*rentalbookdm.RentalBook, error)
	UpdateRentalBook(rentalID int, returnDate time.Time) (*rentalbookdm.RentalBook, error)
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

func (u *rentalBookUsecase) UpdateRentalBook(rentalID int, returnDate time.Time) (*rentalbookdm.RentalBook, error) {
	targetRental, err := u.rentalBookRepository.FetchRentalBook(rentalID)
	if err != nil {
		return nil, err
	}

	if err = targetRental.Set(returnDate); err != nil {
		return nil, err
	}

	updateRentalBook, err := u.rentalBookRepository.UpdateRentalBook(targetRental)
	if err != nil {
		return nil, err
	}

	return updateRentalBook, nil
}

func (u *rentalBookUsecase) FetchRentalBooks() ([]rentalbookdm.RentalBook, error) {
	rentalBooks, err := u.rentalBookRepository.FetchRentalBooks()
	if err != nil {
		return nil, err
	}

	rentalBooksDto := make([]rentalbookdm.RentalBook, len(rentalBooks))

	for i, rentalBook := range rentalBooks {
		rentalBooksDto[i] = rentalbookdm.RentalBook{
			ID:             rentalBook.ID,
			BookID:         rentalBook.BookID,
			UserID:         rentalBook.UserID,
			LoanDate:       rentalBook.LoanDate,
			ReturnDate:     rentalBook.ReturnDate,
			ReturnDeadline: rentalBook.ReturnDeadline,
		}
	}
	return rentalBooksDto, nil
}
