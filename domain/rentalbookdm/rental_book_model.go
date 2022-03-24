package rentalbookdm

import "time"

type RentalBook struct {
	ID             int       `json:"id"`
	BookID         int       `json:"book_id"`
	UserID         int       `json:"user_id"`
	LoanDate       time.Time `json:"loan_date"`
	ReturnDate     time.Time `json:"return_date"`
	ReturnDeadline time.Time `json:"return_deadline"`
}
