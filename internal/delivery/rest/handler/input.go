package handler

type UserInput struct {
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Birthdate string `json:"birthdate"`
}

type UserWithPermissionInput struct {
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Birthdate string `json:"birthdate"`

	Permission string `json:"permission"`
}

type AuthSingInInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
