package main

import (
	"log"
	"os"

	"github.com/glesiosantos/pedidos-mvc-go/internal/routes"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar variaveis de ambiente")
	}

	porta := os.Getenv("API_PORT")

	r := routes.Novo()
	if err := r.Run(":" + porta); err != nil {
		log.Fatal(err)
	}

}
