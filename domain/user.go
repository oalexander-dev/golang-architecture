package domain

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullName"`
}

type UserInput struct {
	Username string `json:"username"`
	FullName string `json:"fullName"`
}

type UserRepo interface {
	GetByID(id int64) (User, error)
	Create(user UserInput) (User, error)
}

type UserOps interface {
	GetByID(id int64) (User, error)
	Create(user UserInput) (User, error)
}
