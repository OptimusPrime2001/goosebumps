package services

import (
	"user-manage-backend/internal/dto"
	"user-manage-backend/internal/models"
)

type UserService interface {
	GetAllUsers(query dto.GetListUserQueryParams) []models.Users
	GetUserByUUID(uuid string) (models.Users, bool)
	CreateUser(user models.Users) (models.Users, error)
	UpdateUser()
	DeleteUser()
}
