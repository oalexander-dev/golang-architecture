package main

import (
	"log"
	"os"

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

	ginMode := os.Getenv("GIN_MODE")
	if ginMode != "debug" {
		ginMode = "release"
	}

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		panic("SECRET_KEY env variable is not defined")
	}
	if len(secretKey) < 36 {
		panic("SECRET_KEY must be at least 36 characters")
	}

	repos := repositories.NewMemoryRepo()
	ops := operations.NewOps(repos)
	app := bindings.NewGinBinding(ops, bindings.GinBindingConfig{
		GinMode:   ginMode,
		SecretKey: secretKey,
	})

	app.Run()
}
