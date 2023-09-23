package main

import (
	"log"

	"github.com/joho/godotenv"
	bindings "github.com/oalexander-dev/golang-architecture/bindings/gin"
	"github.com/oalexander-dev/golang-architecture/operations"
	repositories "github.com/oalexander-dev/golang-architecture/repositories/memoryDB"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	repos := repositories.NewMemoryRepo()
	ops := operations.NewOps(repos)
	app := bindings.NewGinBinding(ops)

	app.Run()
}
