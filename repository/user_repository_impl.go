package repository

import (
	"github.com/nandonyata/Stray-Fedding/config"
	"github.com/nandonyata/Stray-Fedding/entity"
	"github.com/nandonyata/Stray-Fedding/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepositoryImpl{
		Collection: db.Collection("users"),
	}
}

type userRepositoryImpl struct {
	Collection *mongo.Collection
}

func (repo *userRepositoryImpl) Insert(user entity.User) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repo.Collection.InsertOne(ctx, bson.M{
		"_id":      user.Id,
		"email":    user.Email,
		"password": user.Password,
	})
	exception.PanicIfNeeded(err)
}
