package services

import (
	"context"
	"user-manage-backend/internal/db/sqlc"
	"user-manage-backend/internal/dto"
	repositories "user-manage-backend/internal/respositories"
	"user-manage-backend/internal/utils"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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

func (userService *userService) GetAllUsers(ctx context.Context, query dto.GetListUserQueryParams) []sqlc.User {
	var filteredUsers []sqlc.User
	if query.Search != "" {
		filteredUsers = userService.repo.FindBySearch(context.Background(), query.Search)
	} else {
		var err error
		filteredUsers, err = userService.repo.FindAll(context.Background())
		if err != nil {
			return []sqlc.User{}
		}
	}

	start := (query.Page - 1) * query.Limit
	end := start + query.Limit
	if start >= len(filteredUsers) {
		return []sqlc.User{}
	}
	if end > len(filteredUsers) {
		end = len(filteredUsers)
	}
	return filteredUsers[start:end]
}
func (userService *userService) CreateUser(ctx context.Context, user sqlc.CreateUserParams) (sqlc.User, error) {
	user.Email = utils.NormalizeString(user.Email)
	if exist := userService.repo.FindEmail(ctx, user.Email); exist {
		return sqlc.User{}, utils.NewError("email already exists", utils.ErrCodeConflict)
	}
	user.UserID = uuid.New()
	bcryptHash, err := bcrypt.GenerateFromPassword([]byte(user.Password.String), bcrypt.DefaultCost)
	if err != nil {
		return sqlc.User{}, utils.WrapAppError(err, "failed to hash password", utils.ErrCodeInternal)
	}
	user.Password = pgtype.Text{String: string(bcryptHash), Valid: true}
	newUser, err := userService.repo.Create(ctx, user)
	if err != nil {
		return sqlc.User{}, utils.WrapAppError(err, "failed to create user", utils.ErrCodeInternal)
	}
	return newUser, nil
}
func (userService *userService) GetUserByUUID(ctx context.Context, uuid string) (sqlc.User, bool) {
	user, found := userService.repo.FindByUUID(ctx, uuid)
	if !found {
		return sqlc.User{}, false
	}
	return user, true
}
func (userService *userService) UpdateUser(ctx context.Context, uuid string, user sqlc.CreateUserParams) (sqlc.User, error) {
	user.Email = utils.NormalizeString(user.Email)
	if exist := userService.repo.FindEmail(ctx, user.Email); exist {
		return sqlc.User{}, utils.NewError("email already exists", utils.ErrCodeConflict)
	}
	currentUser, found := userService.repo.FindByUUID(ctx, uuid)
	if !found {
		return sqlc.User{}, utils.NewError("user not found", utils.ErrCodeNotFound)
	}
	currentUser.Email = user.Email
	currentUser.Username = user.Username
	currentUser.Age = user.Age
	currentUser.Level = user.Level
	currentUser.Status = user.Status
	if err := userService.repo.Update(ctx, uuid, currentUser); err != nil {
		return sqlc.User{}, utils.WrapAppError(err, "failed to update user", utils.ErrCodeInternal)
	}
	return currentUser, nil
}
func (userService *userService) DeleteUser(ctx context.Context, uuid string) error {
	if err := userService.repo.Delete(ctx, uuid); err != nil {
		return utils.WrapAppError(err, "failed to delete user", utils.ErrCodeInternal)
	}
	return nil
}
