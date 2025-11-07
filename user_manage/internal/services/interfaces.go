package services

import (
	"context"
	"user-manage-backend/internal/db/sqlc"
	"user-manage-backend/internal/dto"
)

type UserService interface {
	GetAllUsers(ctx context.Context, query dto.GetListUserQueryParams) []sqlc.User
	GetUserByUUID(ctx context.Context, uuid string) (sqlc.User, bool)
	CreateUser(ctx context.Context, user sqlc.CreateUserParams) (sqlc.User, error)
	UpdateUser(ctx context.Context, uuid string, user sqlc.CreateUserParams) (sqlc.User, error)
	DeleteUser(ctx context.Context, uuid string) error
}
