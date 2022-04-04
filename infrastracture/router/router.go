package router

import (
	"github.com/labstack/echo"

	"github.com/onituka/go-echo-gorm-sample/infrastracture/persistence"
	"github.com/onituka/go-echo-gorm-sample/infrastracture/persistence/rdb"
	"github.com/onituka/go-echo-gorm-sample/interface/handler"
	"github.com/onituka/go-echo-gorm-sample/usecase"
)

func Run() error {

	bookRepository := persistence.NewBookRepository(rdb.NewDB())
	bookUsecase := usecase.NewBookUsecase(bookRepository)
	bookHandler := handler.NewBookHandler(bookUsecase)

	rentalBookRepository := persistence.NewRentalBookRepository(rdb.NewDB())
	rentalBookUsecase := usecase.NewRentalBookUsecase(rentalBookRepository)
	rentalBookHandler := handler.NewRentalBookHandler(rentalBookUsecase)

	e := echo.New()

	e.GET("/book/:id", bookHandler.Get())
	e.POST("/book", bookHandler.Post())
	e.GET("/books", bookHandler.GetAll())

	e.GET("/rentalbook/:id", rentalBookHandler.Get())
	e.GET("/rentalbook", rentalBookHandler.GetAll())
	e.POST("/rentalbook", rentalBookHandler.Post())
	e.PUT("/rentalbook/:id", rentalBookHandler.Put())

	if err := e.Start(":8015"); err != nil {
		return err
	}

	return nil
}
