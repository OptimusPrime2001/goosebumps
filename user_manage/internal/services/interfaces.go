package services

import "user-manage-backend/internal/models"

type UserService interface {
	GetAllUsers()
	CreateUser(user models.Users) (models.Users, error)
	UpdateUser()
	DeleteUser()
}
