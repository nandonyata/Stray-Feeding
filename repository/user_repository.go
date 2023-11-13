package repository

import "github.com/nandonyata/Stray-Fedding/entity"

type UserRepository interface {
	Insert(user entity.User)
}
