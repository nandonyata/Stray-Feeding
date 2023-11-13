package model

type CreateUserRequest struct {
	email    string `json:"email"`
	password string `json:"password"`
}

type CreateUserResponse struct {
	email    string `json:"email"`
	password string `json:"password"`
}
