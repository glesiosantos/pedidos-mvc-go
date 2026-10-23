package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/glesiosantos/pedidos-mvc-go/internal/database"
	"github.com/glesiosantos/pedidos-mvc-go/internal/routes"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar variaveis de ambiente")
	}

	url := os.Getenv("DATABASE_URL")
	db, err := database.New(context.Background(), url)

	fmt.Println(url)

	if err != nil {
		log.Fatalf("Erro ao configurar/conectar ao PostgreSQL: %v", err)
	}

	db.Close()
	porta := os.Getenv("API_PORT")

	r := routes.Novo()
	if err := r.Run(":" + porta); err != nil {
		log.Fatal(err)
	}
}
