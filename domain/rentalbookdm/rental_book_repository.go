package rentalbookdm

type RentalBookRepository interface {
	FetchRentalBook(rentalID int) (*RentalBook, error)
}
