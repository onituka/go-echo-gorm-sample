package rentalbookdm

type RentalBookRepository interface {
	FetchRentalBook(rentalID int) (*RentalBook, error)
	CreateRentalBook(rentalBook *RentalBook) (*RentalBook, error)
}
