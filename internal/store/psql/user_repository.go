package psql

import (
	"database/sql"

	"github.com/yudgxe/sima-rest-api/internal/model"
	"github.com/yudgxe/sima-rest-api/internal/store"
)

type UserRepository struct {
	db *sql.DB
}

var _ store.UserRepository = (*UserRepository)(nil)

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(u *model.User) error {
	if err := ur.db.QueryRow(
		`INSERT INTO users(
			name,
			surname,
			login,
			password,
			birthdate
		) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		u.Name,
		u.Surname,
		u.Login,
		u.Password,
		u.Birthdate,
	).Scan(&u.ID); err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) UpdateByLogin(u *model.User, login string) error {
	if err := ur.db.QueryRow(
		`UPDATE users SET(
			name,
			surname,
			login,
			password,
			birthdate
		) = ($1, $2, $3, $4, $5) WHERE login = $6 RETURNING id`,
		u.Name,
		u.Surname,
		u.Login,
		u.Password,
		u.Birthdate,
		login,
	).Scan(&u.ID); err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) FindByLogin(login string) (*model.User, error) {
	user := &model.User{}

	if err := ur.db.QueryRow(
		`SELECT id, name, surname, login, password, birthdate FROM users WHERE login = $1`,
		login,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Login,
		&user.Password,
		&user.Birthdate,
	); err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) DeleteByLogin(login string) error {
	tx, err := ur.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var id int

	if err := tx.QueryRow(
		`SELECT id FROM users WHERE login = $1`,
		login,
	).Scan(&id); err != nil {
		return err
	}

	if err := tx.QueryRow(
		`DELETE FROM privileges WHERE user_id = $1 RETURNING user_id`,
		id,
	).Scan(&id); err != nil {
		return err
	}

	if err := tx.QueryRow(
		`DELETE FROM users WHERE login = $1 RETURNING id`,
		login,
	).Scan(&id); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) CreateWithPermission(u *model.User, permission string) error {
	if err := ur.db.QueryRow(
		`INSERT INTO users(
			name,
			surname,
			login,
			password,
			birthdate
		) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		u.Name,
		u.Surname,
		u.Login,
		u.Password,
		u.Birthdate,
	).Scan(&u.ID); err != nil {
		return err
	}

	var userID int
	if err := ur.db.QueryRow(
		`INSERT INTO privileges(user_id, permission) VALUES ($1, $2) RETURNING user_id`,
		u.ID,
		permission,
	).Scan(&userID); err != nil {
		return err
	}

	return nil
}
