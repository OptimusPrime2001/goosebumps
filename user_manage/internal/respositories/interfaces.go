package repositories

import "user-manage-backend/internal/models"

type UserRepository interface {
	FindAll() []models.Users
	FindBySearch(search string) []models.Users
	Create(user models.Users) error
	FindByUUID(uuid string) (models.Users, bool)
	Update()
	Delete()
	FindEmail(email string) bool
}
