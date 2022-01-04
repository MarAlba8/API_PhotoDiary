package models

type Account struct {
	ID             string `json:"id"`
	Nickname       string `json:"nickname"`
	Password       string `json:"password"`
	ProfilePicture string `json:"profilepicture"`
	Username       string `json:"username"`
	Email          string `json:"email"`
}

type Credentials struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type CredentialsToUpdate struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	ProfilePicture string `json:"profilepicture"`
}

type LoginCredentials struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}
