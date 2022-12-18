package mock

import (
	"github.com/stretchr/testify/mock"
	"github.com/yudgxe/sima-rest-api/internal/model"
	"github.com/yudgxe/sima-rest-api/internal/service"
)

type UserService struct {
	mock.Mock
}

var _ service.UserService = (*UserService)(nil)

func (m *UserService) GetUser(login string) (*model.User, error) {
	args := m.Called(login)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *UserService) DeleteUser(login string) error {
	args := m.Called(login)
	return args.Error(0)
}

func (m *UserService) UpdateUser(u *model.User, login string) error {
	args := m.Called(u, login)
	return args.Error(0)
}

func (m *UserService) CreateWithPermission(u *model.User, permission string) error {
	args := m.Called(u, permission)
	return args.Error(0)
}
