package main

import (
	"log"

	"github.com/glesiosantos/pedidos-mvc-go/internal/routes"
)

func main() {

	r := routes.Novo()
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
