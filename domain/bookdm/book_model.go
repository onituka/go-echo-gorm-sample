package bookdm

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func NewBook(id int, title string, author string) *Book {
	return &Book{
		ID:     id,
		Title:  title,
		Author: author,
	}
}
