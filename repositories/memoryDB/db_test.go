package repositories_test

import (
	"testing"

	"github.com/oalexander-dev/golang-architecture/domain"
	repositories "github.com/oalexander-dev/golang-architecture/repositories/memoryDB"
)

func TestNewMemoryRepo(t *testing.T) {
	memoryRepo := repositories.NewMemoryRepo()

	_, err := memoryRepo.User.GetByID(0)

	if err != nil {
		t.Fatal("Something was bad")
	}
}

func TestAddUser(t *testing.T) {
	memoryRepo := repositories.NewMemoryRepo()

	userToAdd := domain.UserInput{
		Username: "testuser1",
		FullName: "Test User 1",
	}

	user, err := memoryRepo.User.Create(userToAdd)
	if err != nil {
		t.Fatal("Unexpected error occurred in user create")
	}

	user, err = memoryRepo.User.GetByID(user.ID)
	if err != nil {
		t.Fatal("Unexpected error occurred retrieving user")
	}

	if user.ID != 1 {
		t.Fatal("Expected user with ID 1, but got something else")
	}
}
