package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/glesiosantos/pedidos-mvc-go/internal/database"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
func run(args []string) error {
	if len(args) == 0 {
		return errors.New("uso: go run ./cmd/migrate status|up|up-to N|version|down")
	}
	arquivo := ".env"
	if len(args) > 1 && args[0] == "-env" {
		arquivo = args[1]
		args = args[2:]
	}
	if len(args) == 0 {
		return errors.New("informe um comando de migration")
	}
	if err := godotenv.Load(arquivo); err != nil {
		return errors.New("arquivo de ambiente ausente; execute na raiz do projeto")
	}
	dsn := os.Getenv("DATABASE_URL")
	if arquivo == ".env.test" {
		dsn = os.Getenv("TEST_DATABASE_URL")
	}
	db, err := database.New(context.Background(), dsn)
	if err != nil {
		return err
	}
	defer db.Close()
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	switch args[0] {
	case "status":
		return goose.Status(db, "migrations")
	case "up":
		return goose.Up(db, "migrations")
	case "up-to":
		if len(args) != 2 {
			return errors.New("informe a versão: up-to 1")
		}
		versao, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil || versao < 1 {
			return errors.New("versão deve ser um inteiro positivo")
		}
		return goose.UpTo(db, "migrations", versao)
	case "version":
		versao, err := goose.GetDBVersion(db)
		if err != nil {
			return err
		}
		fmt.Println(versao)
		return nil
	case "down":
		return goose.Down(db, "migrations")
	default:
		return errors.New("comando de migration desconhecido")
	}
}
