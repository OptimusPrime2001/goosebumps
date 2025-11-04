package repositories

import (
	"log"
	"user-manage-backend/internal/models"
	"user-manage-backend/internal/utils"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
	// TODO: implement repository methods
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (user_repo *userRepository) Create(user models.Users) error {
	// TODO: implement method
	if err := user_repo.db.Create(&user).Error; err != nil {
		return utils.NewError("failed to create user", utils.ErrCodeInternal)
	}
	return nil
}
func (user_repo *userRepository) FindAll() ([]models.Users, error) {
	// TODO: implement method
	var users []models.Users
	if err := user_repo.db.Find(&users).Error; err != nil {
		return users, utils.NewError("failed to find users", utils.ErrCodeInternal)
	}
	return users, nil
}
func (user_repo *userRepository) FindByUUID(uuid string) (models.Users, bool) {
	// TODO: implement method
	var user models.Users
	if err := user_repo.db.Where("user_id = ?", uuid).First(&user).Error; err != nil {
		return models.Users{}, false
	}
	return user, true
}
func (user_repo *userRepository) Update(uuid string, newUser models.Users) error {
	var user models.Users
	if err := user_repo.db.Where("user_id = ?", uuid).First(&user).Error; err != nil {
		return utils.NewError("user not found", utils.ErrCodeNotFound)
	}
	if err := user_repo.db.Model(&user).Updates(newUser).Error; err != nil {
		return utils.NewError("failed to update user", utils.ErrCodeInternal)
	}
	return nil
}
func (user_repo *userRepository) Delete(uuid string) error {
	var user models.Users
	if err := user_repo.db.Where("user_id = ?", uuid).First(&user).Error; err != nil {
		return utils.NewError("user not found", utils.ErrCodeNotFound)
	}
	if err := user_repo.db.Delete(&user).Error; err != nil {
		return utils.NewError("failed to delete user", utils.ErrCodeInternal)
	}
	return nil
}
func (user_repo *userRepository) FindEmail(email string) bool {
	// TODO: implement method
	var user models.Users
	if err := user_repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return false
	}
	return true
}
func (user_repo *userRepository) FindBySearch(search string) []models.Users {
	// TODO: implement method
	if search != "" {
		var filteredUsers []models.Users
		if err := user_repo.db.Where("username LIKE ?", "%"+search+"%").Find(&filteredUsers).Error; err != nil {
			return filteredUsers
		}
		log.Printf("search: %s, filteredUsers: %v", search, filteredUsers)
		return filteredUsers
	} else {
		var users []models.Users
		if err := user_repo.db.Find(&users).Error; err != nil {
			return users
		}
		return users
	}
}
