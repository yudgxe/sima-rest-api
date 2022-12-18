package psql

import (
	"database/sql"

	"github.com/yudgxe/sima-rest-api/internal/model"
	"github.com/yudgxe/sima-rest-api/internal/store"
)

type AuthRepository struct {
	db *sql.DB
}

var _ store.AuthRepository = (*AuthRepository)(nil)

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (ar *AuthRepository) GetUser(login, password string) (*model.Privilege, error) {
	p := &model.Privilege{}

	if err := ar.db.QueryRow(
		`SELECT users.id, privileges.permission 
		FROM users JOIN privileges ON privileges.user_id = users.id 
		WHERE users.login = $1 and users.password = $2`,
		login,
		password,
	).Scan(
		&p.UserID,
		&p.Permission,
	); err != nil {
		return nil, err
	}

	return p, nil
}
