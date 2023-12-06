package service

import (
	"context"

	domain "github.com/lgukasyan/SkyCrypt/domain/user"
	"github.com/lgukasyan/SkyCrypt/internal/infrastructure/security"
	repository "github.com/lgukasyan/SkyCrypt/repository/user"
)

type IUserServiceInterface interface {
	InsertUser(context.Context, *domain.User) error
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
	doc.Password, err = security.Hash(doc.Password)
	
	if err != nil {
		return err
	}

	return us.UserRepository.Save(ctx, doc)
}