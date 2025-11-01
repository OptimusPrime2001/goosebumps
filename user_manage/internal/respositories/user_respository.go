package repositories

import "user-manage-backend/internal/models"

type UserRepository struct {
	users []models.Users
	// TODO: implement repository methods
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make([]models.Users, 0),
	}
}

func (u *UserRepository) CreateUser(user *models.Users) error {
	// TODO: implement method
	return nil
}
