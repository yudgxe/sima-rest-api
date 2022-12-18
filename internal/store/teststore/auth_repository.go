package teststore

import (
	"errors"

	"github.com/yudgxe/sima-rest-api/internal/model"
	"github.com/yudgxe/sima-rest-api/internal/store"
)

type AuthRepository struct {
	users      map[string]*model.User
	privileges map[string]string
}

var _ store.AuthRepository = (*AuthRepository)(nil)

func NewAuthRepository(users map[string]*model.User, privileges map[string]string) *AuthRepository {
	return &AuthRepository{
		users:      users,
		privileges: privileges,
	}
}

func (ar *AuthRepository) GetUser(login, password string) (*model.Privilege, error) {
	u, ok := ar.users[login]
	if !ok {
		return nil, errors.New("no user with this login")
	}

	if u.Password != password {
		return nil, errors.New("no such user")
	}

	p, ok := ar.privileges[login]
	if !ok {
		return nil, errors.New("no privileges with this login")
	}

	return &model.Privilege{
		UserID:     u.ID,
		Permission: p,
	}, nil
}
