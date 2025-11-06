package services

import (
	"context"
	"user-manage-backend/internal/db/sqlc"
	"user-manage-backend/internal/dto"
	"user-manage-backend/internal/models"
	repositories "user-manage-backend/internal/respositories"
	"user-manage-backend/internal/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (userService *userService) GetAllUsers(query dto.GetListUserQueryParams) []models.Users {
	var filteredUsers []models.Users
	if query.Search != "" {
		filteredUsers = userService.repo.FindBySearch(query.Search)
	} else {
		var err error
		filteredUsers, err = userService.repo.FindAll()
		if err != nil {
			return []models.Users{}
		}
	}

	start := (query.Page - 1) * query.Limit
	end := start + query.Limit
	if start >= len(filteredUsers) {
		return []models.Users{}
	}
	if end > len(filteredUsers) {
		end = len(filteredUsers)
	}
	return filteredUsers[start:end]
}
func (userService *userService) CreateUser(ctx context.Context, user sqlc.CreateUserParams) (models.Users, error) {
	user.Email = utils.NormalizeString(user.Email)
	if exist := userService.repo.FindEmail(user.Email); exist {
		return models.Users{}, utils.NewError("email already exists", utils.ErrCodeConflict)
	}
	user.UUID = uuid.NewString()
	bcryptHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.Users{}, utils.WrapAppError(err, "failed to hash password", utils.ErrCodeInternal)
	}
	user.Password = string(bcryptHash)
	if err := userService.repo.Create(user); err != nil {
		return models.Users{}, utils.WrapAppError(err, "failed to create user", utils.ErrCodeInternal)
	}
	return user, nil
}
func (userService *userService) GetUserByUUID(uuid string) (models.Users, bool) {
	user, found := userService.repo.FindByUUID(uuid)
	if !found {
		return models.Users{}, false
	}
	return user, true
}
func (userService *userService) UpdateUser(uuid string, user models.Users) (models.Users, error) {
	user.Email = utils.NormalizeString(user.Email)
	if exist := userService.repo.FindEmail(user.Email); exist {
		return models.Users{}, utils.NewError("email already exists", utils.ErrCodeConflict)
	}
	currentUser, found := userService.repo.FindByUUID(uuid)
	if !found {
		return models.Users{}, utils.NewError("user not found", utils.ErrCodeNotFound)
	}
	currentUser.Email = user.Email
	currentUser.Username = user.Username
	currentUser.Age = user.Age
	currentUser.Level = user.Level
	currentUser.Status = user.Status
	if err := userService.repo.Update(uuid, currentUser); err != nil {
		return models.Users{}, utils.WrapAppError(err, "failed to update user", utils.ErrCodeInternal)
	}
	return currentUser, nil
}
func (userService *userService) DeleteUser(uuid string) error {
	if err := userService.repo.Delete(uuid); err != nil {
		return utils.WrapAppError(err, "failed to delete user", utils.ErrCodeInternal)
	}
	return nil
}
