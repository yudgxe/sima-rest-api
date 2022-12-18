package teststore

import (
	"github.com/yudgxe/sima-rest-api/internal/model"
	"github.com/yudgxe/sima-rest-api/internal/store"
)

type Store struct {
	userRepository *UserRepository
	authRepository *AuthRepository
}

var _ store.Store = (*Store)(nil)

func New() *Store {
	users := make(map[string]*model.User)
	privileges := make(map[string]string)

	return &Store{
		userRepository: NewUserRepository(users, privileges),
		authRepository: NewAuthRepository(users, privileges),
	}
}

func (s *Store) User() store.UserRepository {
	return s.userRepository
}

func (s *Store) Auth() store.AuthRepository {
	return s.authRepository
}
