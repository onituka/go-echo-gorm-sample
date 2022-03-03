package bookdm

type BookRepository interface {
	FetchBook(BookID int) (*Book, error)
}
