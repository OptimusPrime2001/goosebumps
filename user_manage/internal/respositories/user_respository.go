package repositories

import "user-manage-backend/internal/models"

type userRepository struct {
	users []models.Users
	// TODO: implement repository methods
}

func NewUserRepository() UserRepository {
	return &userRepository{
		users: make([]models.Users, 0),
	}
}

func (user_repo *userRepository) Create(user models.Users) error {
	// TODO: implement method
	user_repo.users = append(user_repo.users, user)
	return nil
}
func (user_repo *userRepository) FindAll() {
	// TODO: implement method
}
func (user_repo *userRepository) FindByUUID() {
	// TODO: implement method
}
func (user_repo *userRepository) Update() {
	// TODO: implement method
}
func (user_repo *userRepository) Delete() {
	// TODO: implement method
}
func (user_repo *userRepository) FindEmail(email string) bool {
	// TODO: implement method
	for _, user := range user_repo.users {
		if user.Email == email {
			return true
		}
	}
	return false
}
