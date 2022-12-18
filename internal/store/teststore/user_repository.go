package teststore

import (
	"errors"

	"github.com/yudgxe/sima-rest-api/internal/model"
	"github.com/yudgxe/sima-rest-api/internal/store"
)

type UserRepository struct {
	users      map[string]*model.User
	privileges map[string]string
}

var _ store.UserRepository = (*UserRepository)(nil)

func NewUserRepository(users map[string]*model.User, privileges map[string]string) *UserRepository {
	return &UserRepository{
		users:      users,
		privileges: privileges,
	}
}

func (ur *UserRepository) Create(u *model.User) error {
	u.ID = (len(ur.users) + 1)
	ur.users[u.Login] = u

	return nil
}

func (ur *UserRepository) UpdateByLogin(u *model.User, login string) error {
	oldU, ok := ur.users[u.Login]
	if !ok {
		return errors.New("no user with this login")
	}

	u.ID = oldU.ID
	ur.users[u.Login] = u

	return nil
}

func (ur *UserRepository) FindByLogin(login string) (*model.User, error) {
	u, ok := ur.users[login]
	if !ok {
		return nil, errors.New("no user with this login")
	}

	return u, nil
}

func (ur *UserRepository) DeleteByLogin(login string) error {
	delete(ur.users, login)
	return nil
}

func (ur *UserRepository) CreateWithPermission(u *model.User, permission string) error {
	u.ID = (len(ur.users) + 1)
	ur.users[u.Login] = u
	ur.privileges[u.Login] = permission

	return nil
}
