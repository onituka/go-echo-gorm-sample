package handler

import (
	"github.com/labstack/echo"
	"github.com/onituka/go-echo-gorm-sample/usecase"
	"net/http"
	"strconv"
)

type BookHandler interface {
	Get() echo.HandlerFunc
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
