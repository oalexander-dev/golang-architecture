package repositories_test

import (
	"testing"

	repositories "github.com/oalexander-dev/golang-architecture/repositories/memoryDB"
)

func TestNewMemoryRepo(t *testing.T) {
	repositories.NewMemoryRepo()
}
