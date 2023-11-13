package service

import "github.com/nandonyata/Stray-Fedding/model"

type UserService interface {
	Create(request model.CreateUserRequest) (response model.CreateUserResponse)
}
