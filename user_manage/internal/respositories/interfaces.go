package repositories

import (
	"context"
	"user-manage-backend/internal/db/sqlc"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]sqlc.User, error)
	FindBySearch(ctx context.Context, search string) []sqlc.User
	Create(ctx context.Context, user sqlc.CreateUserParams) (sqlc.User, error)
	FindByUUID(ctx context.Context, uuid string) (sqlc.User, bool)
	Update(ctx context.Context, uuid string, user sqlc.User) error
	Delete(ctx context.Context, uuid string) error
	FindEmail(ctx context.Context, email string) bool
}
