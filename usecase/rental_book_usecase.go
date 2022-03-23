package usecase

import "github.com/onituka/go-echo-gorm-sample/domain/rentalbookdm"

type RentalBookUsecase interface {
	FetchRentalBook(rentalID int) (*rentalbookdm.RentalBook, error)
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
