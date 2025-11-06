package repositories

import (
	"context"
	"log"
	"user-manage-backend/internal/db/sqlc"
	"user-manage-backend/internal/models"
	"user-manage-backend/internal/utils"
)

type userRepository struct {
	db sqlc.Querier
	// TODO: implement repository methods
}

func NewUserRepository(db sqlc.Querier) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (user_repo *userRepository) Create(ctx context.Context, input sqlc.CreateUserParams) (sqlc.User, error) {
	// TODO: implement method
	user, err := user_repo.db.CreateUser(ctx, input)
	if err != nil {
		return sqlc.User{}, utils.NewError("failed to create user", utils.ErrCodeInternal)
	}
	return user, nil
}
func (user_repo *userRepository) FindAll(ctx context.Context) ([]models.Users, error) {
	// TODO: implement method
	var users []models.Users
	if err := user_repo.db.WithContext(ctx).Find(&users).Error; err != nil {
		return users, utils.NewError("failed to find users", utils.ErrCodeInternal)
	}
	return users, nil
}
func (user_repo *userRepository) FindByUUID(ctx context.Context, uuid string) (models.Users, bool) {
	// TODO: implement method
	var user models.Users
	if err := user_repo.db.WithContext(ctx).Where("user_id = ?", uuid).First(&user).Error; err != nil {
		return models.Users{}, false
	}
	return user, true
}
func (user_repo *userRepository) Update(ctx context.Context, uuid string, newUser models.Users) error {
	var user models.Users
	if err := user_repo.db.WithContext(ctx).Where("user_id = ?", uuid).First(&user).Error; err != nil {
		return utils.NewError("user not found", utils.ErrCodeNotFound)
	}
	if err := user_repo.db.WithContext(ctx).Model(&user).Updates(newUser).Error; err != nil {
		return utils.NewError("failed to update user", utils.ErrCodeInternal)
	}
	return nil
}
func (user_repo *userRepository) Delete(ctx context.Context, uuid string) error {
	var user models.Users
	if err := user_repo.db.WithContext(ctx).Where("user_id = ?", uuid).First(&user).Error; err != nil {
		return utils.NewError("user not found", utils.ErrCodeNotFound)
	}
	if err := user_repo.db.WithContext(ctx).Delete(&user).Error; err != nil {
		return utils.NewError("failed to delete user", utils.ErrCodeInternal)
	}
	return nil
}
func (user_repo *userRepository) FindEmail(ctx context.Context, email string) bool {
	// TODO: implement method
	var user models.Users
	if err := user_repo.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return false
	}
	return true
}
func (user_repo *userRepository) FindBySearch(ctx context.Context, search string) []models.Users {
	// TODO: implement method
	if search != "" {
		var filteredUsers []models.Users
		if err := user_repo.db.WithContext(ctx).Where("username LIKE ?", "%"+search+"%").Find(&filteredUsers).Error; err != nil {
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
