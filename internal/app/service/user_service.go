package service

import (
	"context"
	"errors"

	domain "github.com/lgukasyan/SkyCrypt/domain/user"
	"github.com/lgukasyan/SkyCrypt/internal/infrastructure/security"
	repository "github.com/lgukasyan/SkyCrypt/repository/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserServiceInterface interface {
	InsertUser(context.Context, *domain.User) error
	FindAndValidate(context.Context, *domain.UserSignIn) (*domain.User, error)
}

type UserServiceImpl struct {
	UserRepository repository.IUserRepositoryInterface
}

func NewUserService(userRepo repository.IUserRepositoryInterface) IUserServiceInterface {
	return &UserServiceImpl{
		UserRepository: userRepo,
	}
}

func (us *UserServiceImpl) InsertUser(ctx context.Context, doc *domain.User) error {
	var err error
	(*doc).Password, err = security.Hash((*doc).Password)
	(*doc).Id = primitive.NewObjectID()

	if err != nil {
		return err
	}

	return us.UserRepository.Save(ctx, doc)
}

func (us *UserServiceImpl) FindAndValidate(ctx context.Context, user *domain.UserSignIn) (*domain.User, error) {
	doc, err := us.UserRepository.FindByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if !security.CheckPasswordHash(user.Password, doc.Password) {
		return nil, errors.New("incorrect e-mail address or password")
	}

	return doc, nil
}
