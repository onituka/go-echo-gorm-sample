package bookdm

import "errors"

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Number string `json:"number"`
}

func NewBook(title string, author string) (*Book, error) {
	if title == " " || author == "" {
		return nil, errors.New("タイトルもしくは著者名を入力してください")
	}
	return &Book{
		Title:  title,
		Author: author,
	}, nil
}
