package basic

import (
	"github.com/yudgxe/sima-rest-api/internal/model"
	"github.com/yudgxe/sima-rest-api/internal/service"
	"github.com/yudgxe/sima-rest-api/internal/store"
)

type UserService struct {
	store store.Store
}

var _ service.UserService = (*UserService)(nil)

func NewUserService(s store.Store) *UserService {
	return &UserService{
		store: s,
	}
}

func (us *UserService) GetUser(login string) (*model.User, error) {
	return us.store.User().FindByLogin(login)
}

func (us *UserService) DeleteUser(login string) error {
	return us.store.User().DeleteByLogin(login)
}

func (us *UserService) UpdateUser(u *model.User, login string) error {
	return us.store.User().UpdateByLogin(u, login)
}

func (us *UserService) CreateWithPermission(u *model.User, permission string) error {
	return us.store.User().CreateWithPermission(u, permission)
}
