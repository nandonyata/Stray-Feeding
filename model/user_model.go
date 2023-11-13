package model

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
