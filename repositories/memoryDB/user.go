package repositories

import "github.com/oalexander-dev/golang-architecture/domain"

type userRepo struct {
	DB *data
}

func newUserRepo(db *data) userRepo {
	return userRepo{
		DB: db,
	}
}

func (r userRepo) GetByID(id int64) (domain.User, error) {
	for _, user := range r.DB.Users {
		if user.ID == id {
			return user, nil
		}
	}

	return domain.User{}, nil
}

func (r userRepo) Create(user domain.User) (domain.User, error) {
	r.DB.Users = append(r.DB.Users, user)

	return user, nil
}
