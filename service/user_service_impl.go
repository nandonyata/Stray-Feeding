package service

import (
	"github.com/nandonyata/Stray-Fedding/entity"
	"github.com/nandonyata/Stray-Fedding/model"
	"github.com/nandonyata/Stray-Fedding/repository"
)

func NewUserService(userRepo *repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: *userRepo,
	}
}

type userServiceImpl struct {
	UserRepository repository.UserRepository
}

func (service *userServiceImpl) Create(request model.CreateUserRequest) (response model.CreateUserResponse) {
	user := entity.User{
		Id:       request.Id,
		Email:    request.Email,
		Password: request.Password,
	}

	service.UserRepository.Insert(user)

	response = model.CreateUserResponse{
		Id:       user.Id,
		Email:    user.Email,
		Password: user.Password,
	}

	return response
}
