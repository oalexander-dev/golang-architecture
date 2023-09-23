package domain

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullName"`
}

type UserRepo interface {
	GetByID(id int64) (User, error)
	Create(user User) (User, error)
}

type UserOps interface {
	GetByID(id int64) (User, error)
	Create(user User) (User, error)
}
