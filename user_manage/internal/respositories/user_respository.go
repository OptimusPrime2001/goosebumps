package repositories

import (
	"log"
	"strings"
	"user-manage-backend/internal/models"
	"user-manage-backend/internal/utils"

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
func (user_repo *userRepository) Update(uuid string, newUser models.Users) error {
	// TODO: implement method
	_, index, ok := lo.FindIndexOf(user_repo.users, func(u models.Users) bool {
		return u.UUID == uuid
	})
	if !ok {
		return utils.NewError("user not found", utils.ErrCodeNotFound)
	}
	user_repo.users[index] = newUser
	return nil
}
func (user_repo *userRepository) Delete(uuid string) error {
	// TODO: implement method
	for i, u := range user_repo.users {
		if u.UUID == uuid {
			user_repo.users = append(user_repo.users[:i], user_repo.users[i+1:]...)
			return nil
		}
	}
	return utils.NewError("user not found", utils.ErrCodeNotFound)
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
