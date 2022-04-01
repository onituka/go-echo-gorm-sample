package persistence

import (
	"github.com/jinzhu/gorm"

	"github.com/onituka/go-echo-gorm-sample/domain/rentalbookdm"
)

type rentalBookRepository struct {
	Conn *gorm.DB
}

func NewRentalBookRepository(conn *gorm.DB) rentalbookdm.RentalBookRepository {
	return &rentalBookRepository{Conn: conn}
}

func (r *rentalBookRepository) FetchRentalBook(rentalID int) (*rentalbookdm.RentalBook, error) {
	rentalBook := &rentalbookdm.RentalBook{ID: rentalID}

	if err := r.Conn.Debug().First(&rentalBook).Error; err != nil {
		return nil, err
	}

	return rentalBook, nil
}

func (r *rentalBookRepository) CreateRentalBook(rentalBook *rentalbookdm.RentalBook) (*rentalbookdm.RentalBook, error) {
	if err := r.Conn.Debug().Create(&rentalBook).Error; err != nil {
		return nil, err
	}

	return rentalBook, nil
}

func (r *rentalBookRepository) UpdateRentalBook(rentalBook *rentalbookdm.RentalBook) (*rentalbookdm.RentalBook, error) {
	if err := r.Conn.Model(&rentalBook).Debug().Update(&rentalBook).Error; err != nil {
		return nil, err
	}

	return rentalBook, nil
}
