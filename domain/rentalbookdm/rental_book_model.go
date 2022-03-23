package rentalbookdm

import "time"

type RentalBook struct {
	ID             int
	BookID         int
	UserID         int
	LoanDate       time.Time
	ReturnDate     time.Time
	ReturnDeadline time.Time
}
