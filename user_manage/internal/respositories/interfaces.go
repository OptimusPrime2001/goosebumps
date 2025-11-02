package repositories

import "user-manage-backend/internal/models"

type UserRepository interface {
	FindAll()
	Create(user models.Users) error
	FindByUUID()
	Update()
	Delete()
	FindEmail(email string) bool
}
