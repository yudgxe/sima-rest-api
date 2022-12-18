package mock

import (
	"github.com/stretchr/testify/mock"
	"github.com/yudgxe/sima-rest-api/internal/service"
)

type AuthService struct {
	mock.Mock
}

var _ service.AuthService = (*AuthService)(nil)

func (m *AuthService) GenerateToken(login, password string) (string, error) {
	args := m.Called(login, password)

	return args.Get(0).(string), args.Error(1)
}
