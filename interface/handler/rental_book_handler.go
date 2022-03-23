package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/onituka/go-echo-gorm-sample/usecase"
)

type RentalBookHandler interface {
	Get() echo.HandlerFunc
}

type rentalBookHandler struct {
	rentalBookUsecase usecase.RentalBookUsecase
}

func NewRentalBookHandler(rentalBookUsecase usecase.RentalBookUsecase) *rentalBookHandler {
	return &rentalBookHandler{rentalBookUsecase: rentalBookUsecase}
}

func (h *rentalBookHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		rentalID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		out, err := h.rentalBookUsecase.FetchRentalBook(rentalID)
		if err != nil {
			return c.JSON(http.StatusNoContent, err.Error())
		}
		return c.JSON(http.StatusOK, out)
	}
}
