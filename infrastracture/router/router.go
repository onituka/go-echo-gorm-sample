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

	e := echo.New()

	e.GET("/book/:id", bookHandler.Get())

	if err := e.Start(":8015"); err != nil {
		return err
	}

	return nil
}
