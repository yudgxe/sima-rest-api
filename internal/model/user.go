package model

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Login     string    `json:"login"`
	Password  string    `json:"password,omitempty"`
	Birthdate time.Time `json:"birthdate"`
}

func (u *User) Sanitize() {
	u.Password = ""
}
