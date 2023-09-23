package repositories

import "github.com/oalexander-dev/golang-architecture/domain"

type data struct {
	Users []domain.User
}

func NewMemoryRepo() domain.Repo {
	db := &data{
		Users: make([]domain.User, 0),
	}

	repos := domain.Repo{
		User: newUserRepo(db),
	}

	return repos
}
