package repositories_test

import (
	"testing"

	"github.com/oalexander-dev/golang-architecture/domain"
	repositories "github.com/oalexander-dev/golang-architecture/repositories/memoryDB"
)

var validUserInputs = []domain.UserInput{
	{
		Username: "testuser1",
		FullName: "Test User 1",
	},
	{
		Username: "testuser2",
		FullName: "Test User 2",
	},
	{
		Username: "testuser3",
		FullName: "Test User 3",
	},
}

func addValidUsers(r domain.Repo, t *testing.T) {
	for _, user := range validUserInputs {
		_, err := r.User.Create(user)
		if err != nil {
			t.Fatal("Failed to create one of the valid users")
		}
	}
}

func TestAddUser(t *testing.T) {
	memoryRepo := repositories.NewMemoryRepo()

	userToAdd := domain.UserInput{
		Username: "testuser1",
		FullName: "Test User 1",
	}

	originalUser, err := memoryRepo.User.Create(userToAdd)
	if err != nil {
		t.Fatal("Unexpected error occurred in user create")
	}

	user, err := memoryRepo.User.GetByID(originalUser.ID)
	if err != nil {
		t.Fatal("Unexpected error occurred retrieving user")
	}

	if user.ID != originalUser.ID || user.Username != userToAdd.Username || user.FullName != userToAdd.FullName {
		t.Fatal("Wrong value for a field in the saved user")
	}
}

func TestAddUserMultiple(t *testing.T) {
	memoryRepo := repositories.NewMemoryRepo()

	userToAdd := domain.UserInput{
		Username: "testuser1",
		FullName: "Test User 1",
	}

	originalUser, err := memoryRepo.User.Create(userToAdd)
	if err != nil {
		t.Fatal("Unexpected error occurred in user create")
	}

	secondUser, err := memoryRepo.User.Create(userToAdd)
	if err != nil {
		t.Fatal("Unexpected error occurred in user create")
	}

	user, err := memoryRepo.User.GetByID(originalUser.ID)
	if err != nil {
		t.Fatal("Unexpected error occurred retrieving first user")
	}

	if user.ID != originalUser.ID || user.Username != userToAdd.Username || user.FullName != userToAdd.FullName {
		t.Fatal("Wrong value for a field in the saved user")
	}

	user, err = memoryRepo.User.GetByID(secondUser.ID)
	if err != nil {
		t.Fatal("Unexpected error occurred retrieving second user")
	}

	if user.ID != secondUser.ID || user.Username != userToAdd.Username || user.FullName != userToAdd.FullName {
		t.Fatal("Wrong value for a field in the saved user")
	}
}

func TestGetUser(t *testing.T) {
	memoryRepo := repositories.NewMemoryRepo()

	userToAdd := validUserInputs[0]

	originalUser, err := memoryRepo.User.Create(userToAdd)
	if err != nil {
		t.Fatal("Unexpected error occurred in user create")
	}

	user, err := memoryRepo.User.GetByID(originalUser.ID)
	if err != nil {
		t.Fatal("Unexpected error occurred retrieving user")
	}

	if user.ID != originalUser.ID || user.Username != userToAdd.Username || user.FullName != userToAdd.FullName {
		t.Fatal("Wrong value for a field in the saved user")
	}
}

func TestGetUserWithMultiple(t *testing.T) {
	memoryRepo := repositories.NewMemoryRepo()

	addValidUsers(memoryRepo, t)

	userToAdd := domain.UserInput{
		Username: "testuser4",
		FullName: "Test User 4",
	}

	originalUser, err := memoryRepo.User.Create(userToAdd)
	if err != nil {
		t.Fatal("Unexpected error occurred in user create")
	}

	user, err := memoryRepo.User.GetByID(originalUser.ID)
	if err != nil {
		t.Fatal("Unexpected error occurred retrieving user")
	}

	if user.ID != originalUser.ID || user.Username != userToAdd.Username || user.FullName != userToAdd.FullName {
		t.Fatal("Wrong value for a field in the saved user")
	}
}
