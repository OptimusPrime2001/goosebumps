package services

import (
	"context"
	"user-manage-backend/internal/db/sqlc"
	"user-manage-backend/internal/dto"
	"user-manage-backend/internal/models"
)

type UserService interface {
	GetAllUsers(query dto.GetListUserQueryParams) []models.Users
	GetUserByUUID(uuid string) (models.Users, bool)
	CreateUser(ctx context.Context, user sqlc.CreateUserParams) (models.Users, error)
	UpdateUser(ctx context.Context, uuid string, user models.Users) (models.Users, error)
	DeleteUser(ctx context.Context, uuid string) error
}
