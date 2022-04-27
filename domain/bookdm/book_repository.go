package bookdm

type BookRepository interface {
	FetchBook(bookID int) (*Book, error)
	CreateBook(book *Book) (*Book, error)
	FetchBooks() ([]Book, error)
	SearchBooks(title string, author string) ([]Book, error)
}
