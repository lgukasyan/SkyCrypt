package repository

import (
	"context"
	"errors"

	domain "github.com/lgukasyan/SkyCrypt/domain/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(Coll *mongo.Collection) IUserRepositoryInterface {
	return &UserRepository{
		Collection: Coll,
	}
}

func (ur *UserRepository) Save(ctx context.Context, user *domain.User) error {
	_, err := ur.Collection.InsertOne(ctx, *user);

	if mongo.IsDuplicateKeyError(err) {
		return errors.New("email is already in use")
	}

	return err
}
