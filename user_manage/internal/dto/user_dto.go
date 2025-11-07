package dto

import (
	"user-manage-backend/internal/db/sqlc"
)

// --- Param struct ---
type UserParam struct {
	UUID string `uri:"uuid" binding:"uuid"`
}

// --- Query struct ---
type GetListUserQueryParams struct {
	Search string `form:"search" binding:"omitempty,min=3,max=50,search"`
	Page   int    `form:"page,default=1" binding:"required,gte=1,lte=100"`
	Limit  int    `form:"limit,default=10" binding:"omitempty,gte=1,lte=100"`
}

type UserDTO struct {
	UUID     string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Status   string `json:"status"`
	Level    string `json:"level"`
}

func MapUserToDTO(user sqlc.User) *UserDTO {
	return &UserDTO{
		UUID:     user.UserID.String(),
		Username: user.Username,
		Email:    user.Email,
		Age:      int(user.Age.Int32),
		Status:   MapStatusText(int(user.Status.Int32)),
		Level:    MapLevelText(int(user.Level.Int32)),
	}
}
func MapUsersToDTOs(users []sqlc.User) []*UserDTO {
	userDTOs := make([]*UserDTO, 0)
	for _, user := range users {
		userDTOs = append(userDTOs, MapUserToDTO(user))
	}
	return userDTOs
}
func MapStatusText(status int) string {
	switch status {
	case 1:
		return "active"
	case 0:
		return "inactive"
	default:
		return "unknown"
	}
}
func MapLevelText(level int) string {
	switch level {
	case 1:
		return "user"
	case 0:
		return "admin"
	default:
		return "mode"
	}
}
