package domain

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Password string `json:"password"`
}

type UserInput struct {
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Password string `json:"password"`
}

type UserLoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRepo interface {
	GetByID(id int64) (User, error)
	GetByUsername(username string) (User, error)
	Create(user UserInput) (User, error)
}

type UserOps interface {
	GetByID(id int64) (User, error)
	Create(user UserInput) (User, error)
	CheckPassword(username string, passwordInput string) (User, error)
}
