package operations

import "github.com/oalexander-dev/golang-architecture/domain"

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

func (u userOps) Create(user domain.User) (domain.User, error) {
	user, err := u.Repo.User.Create(user)
	return user, err
}
