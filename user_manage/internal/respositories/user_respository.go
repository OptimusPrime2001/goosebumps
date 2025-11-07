package repositories

import (
	"context"
	"user-manage-backend/internal/db/sqlc"
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
func (user_repo *userRepository) FindAll(ctx context.Context) ([]sqlc.User, error) {
	// TODO: implement method

	return []sqlc.User{}, nil
}
func (user_repo *userRepository) FindByUUID(ctx context.Context, uuid string) (sqlc.User, bool) {
	// TODO: implement method

	return sqlc.User{}, true
}
func (user_repo *userRepository) Update(ctx context.Context, uuid string, newUser sqlc.User) error {
	// TODO: implement method

	return nil
}
func (user_repo *userRepository) Delete(ctx context.Context, uuid string) error {

	return nil
}
func (user_repo *userRepository) FindEmail(ctx context.Context, email string) bool {

	return true
}
func (user_repo *userRepository) FindBySearch(ctx context.Context, search string) []sqlc.User {
	// TODO: implement method
	return []sqlc.User{}
}
