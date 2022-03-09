package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/onituka/go-echo-gorm-sample/domain/bookdm"
	"github.com/onituka/go-echo-gorm-sample/usecase"
)

type BookHandler interface {
	Get() echo.HandlerFunc
	Post() echo.HandlerFunc
	GetAll() echo.HandlerFunc
}

type bookHandler struct {
	bookUsecase usecase.BookUsecase
}

func NewBookHandler(bookUsecase usecase.BookUsecase) BookHandler {
	return &bookHandler{bookUsecase: bookUsecase}
}

func (h *bookHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		bookID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		out, err := h.bookUsecase.FetchBook(bookID)
		if err != nil {
			return c.JSON(http.StatusNoContent, err.Error())
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (h *bookHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req bookdm.Book
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createBook, err := h.bookUsecase.CreateBook(req.Title, req.Author)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		out := bookdm.Book{
			ID:     createBook.ID,
			Title:  createBook.Title,
			Author: createBook.Author,
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (h *bookHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		books, err := h.bookUsecase.FetchBooks()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, books)
	}
}
