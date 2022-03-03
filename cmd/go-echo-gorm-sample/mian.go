package main

import (
	"log"

	"github.com/onituka/go-echo-gorm-sample/infrastracture/router"
)

func main() {
	if err := router.Run(); err != nil {
		log.Fatalln(err)
	}
}
