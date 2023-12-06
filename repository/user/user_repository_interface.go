package repository

import (
	"context"

	domain "github.com/lgukasyan/SkyCrypt/domain/user"
)

type IUserRepositoryInterface interface {
	Save(ctx context.Context, doc *domain.User) error
}