package service

import "github.com/yudgxe/sima-rest-api/internal/model"

type UserService interface {
	GetUser(login string) (*model.User, error)
	DeleteUser(login string) error
	UpdateUser(u *model.User, login string) error
	CreateWithPermission(u *model.User, permission string) error
}

type AuthService interface {
	GenerateToken(login, password string) (string, error)
}
