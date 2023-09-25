package domain

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Password string `json:"-"`
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
	GetByUsername(username string) (User, error)
	Create(user UserInput) (User, error)
}

type UserOps interface {
	GetByUsername(username string) (User, error)
	Create(user UserInput) (User, error)
	CheckPassword(username string, passwordInput string) (User, error)
}
