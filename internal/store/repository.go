package store

import (
	"github.com/yudgxe/sima-rest-api/internal/model"
)

type UserRepository interface {
	Create(u *model.User) error
	UpdateByLogin(u *model.User, login string) error
	FindByLogin(login string) (*model.User, error)
	DeleteByLogin(login string) error
	CreateWithPermission(u *model.User, permission string) error
}

type AuthRepository interface {
	GetUser(login, password string) (*model.Privilege, error)
}
