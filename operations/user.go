package operations

import (
	"errors"

	"github.com/oalexander-dev/golang-architecture/domain"
	"golang.org/x/crypto/bcrypt"
)

type userOps struct {
	Repo domain.Repo
}

func newUserOps(r domain.Repo) domain.UserOps {
	return userOps{
		Repo: r,
	}
}

func (u userOps) GetByID(id int64) (domain.User, error) {
	user, err := u.Repo.User.GetByID(id)
	return user, err
}

func (u userOps) Create(user domain.UserInput) (domain.User, error) {
	_, err := u.Repo.User.GetByUsername(user.Username)
	if err == nil {
		return domain.User{}, errors.New("already exists")
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 15)
	if err != nil {
		return domain.User{}, err
	}

	user.Password = string(hashedBytes)

	savedUser, err := u.Repo.User.Create(user)
	return savedUser, err
}

func (u userOps) CheckPassword(username string, passwordInput string) (domain.User, error) {
	user, err := u.Repo.User.GetByUsername(username)
	if err != nil {
		return domain.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordInput))
	if err != nil {
		return domain.User{}, errors.New("bad credentials")
	}

	return user, nil
}
