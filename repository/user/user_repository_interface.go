package repository

import (
	"context"
	domain "github.com/lgukasyan/SkyCrypt/domain/user"
)

type IUserRepositoryInterface interface {
	Save(ctx context.Context, doc *domain.User) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}