package repositories

import (
	"log"
	"strings"
	"user-manage-backend/internal/models"

	"github.com/samber/lo"
)

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
func (user_repo *userRepository) FindAll() []models.Users {
	// TODO: implement method
	return user_repo.users
}
func (user_repo *userRepository) FindByUUID(uuid string) (models.Users, bool) {
	// TODO: implement method
	for _, user := range user_repo.users {
		if user.UUID == uuid {
			return user, true
		}
	}
	return models.Users{}, false
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
func (user_repo *userRepository) FindBySearch(search string) []models.Users {
	// TODO: implement method
	if search != "" {
		filteredUsers := lo.Filter(user_repo.users, func(user models.Users, index int) bool {
			return strings.Contains(user.Username, search)
		})
		log.Printf("search: %s,users: %v, filteredUsers: %v", search, user_repo.users, filteredUsers)
		return filteredUsers
	} else {
		return user_repo.users
	}
}
