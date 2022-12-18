package psql

import (
	"database/sql"

	"github.com/yudgxe/sima-rest-api/internal/store"
)

type Store struct {
	userRepository *UserRepository
	authRepository *AuthRepository
}

var _ store.Store = (*Store)(nil)

func New(db *sql.DB) *Store {
	return &Store{
		userRepository: NewUserRepository(db),
		authRepository: NewAuthRepository(db),
	}
}

func (s *Store) User() store.UserRepository {
	return s.userRepository
}

func (s *Store) Auth() store.AuthRepository {
	return s.authRepository
}
