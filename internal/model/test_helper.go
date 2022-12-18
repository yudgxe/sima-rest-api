package model

import (
	"testing"
	"time"
)

func GetTestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		Name:      "name",
		Surname:   "surname",
		Login:     "login",
		Password:  "password",
		Birthdate: time.Date(1970, time.Month(1), 1, 0, 0, 0, 0, time.UTC),
	}
}
