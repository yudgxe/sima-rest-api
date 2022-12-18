package database

import (
	"database/sql"
	"fmt"
)

func NewPostgres(info PostgresConnInfo) (*sql.DB, error) {
	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		info.Host,
		info.Port,
		info.User,
		info.Password,
		info.Name,
		info.SSLMode,
	)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
