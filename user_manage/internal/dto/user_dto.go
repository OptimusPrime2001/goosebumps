package dto

import "user-manage-backend/internal/models"

// --- Param struct ---
type UserParam struct {
	UUID string `uri:"id" binding:"uuid"`
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

func MapUserToDTO(user models.Users) *UserDTO {
	return &UserDTO{
		UUID:     user.UUID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
		Status:   MapStatusText(user.Status),
		Level:    MapLevelText(user.Level),
	}
}
func MapUsersToDTOs(users []models.Users) []*UserDTO {
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
