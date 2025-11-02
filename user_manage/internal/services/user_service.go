package services

import (
	"log"
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
	log.Printf("query: %v", query)
	if query.Search != "" {
		return userService.repo.FindBySearch(query.Search)
	}
	return userService.repo.FindAll()
}
func (userService *userService) CreateUser(user models.Users) (models.Users, error) {
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
func (userService *userService) UpdateUser() {
}
func (userService *userService) DeleteUser() {
}
