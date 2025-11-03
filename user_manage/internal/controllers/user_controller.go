package controllers

import (
	"fmt"
	"net/http"
	"user-manage-backend/internal/dto"
	"user-manage-backend/internal/models"
	"user-manage-backend/internal/services"
	"user-manage-backend/internal/utils"
	"user-manage-backend/internal/validations"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (userController *UserController) CreateNewUser(ctx *gin.Context) {
	var user models.Users
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.ResponseValidator(ctx, validations.HandleValidationErrors(err))
		return
	}
	user, err := userController.service.CreateUser(user)
	if err != nil {
		fmt.Println("error:", err)
		utils.ResponseError(ctx, err)
		return
	}
	userDTO := dto.MapUserToDTO(user)
	utils.ResponseSuccess(ctx, http.StatusCreated, &userDTO)
}
func (userController *UserController) GetUserByUUID(ctx *gin.Context) {
	var params dto.UserParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validations.HandleValidationErrors(err))
		return
	}
	user, err := userController.service.GetUserByUUID(params.UUID)
	if !err {
		utils.ResponseError(ctx, utils.NewError("user not found", utils.ErrCodeNotFound))
		return
	}
	userDTO := dto.MapUserToDTO(user)
	utils.ResponseSuccess(ctx, http.StatusOK, &userDTO)
}

func (userController *UserController) GetAllUsers(ctx *gin.Context) {
	var query dto.GetListUserQueryParams
	if err := ctx.ShouldBindQuery(&query); err != nil {
		utils.ResponseValidator(ctx, validations.HandleValidationErrors(err))
		return
	}
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.Limit <= 0 {
		query.Limit = 5
	}
	users := userController.service.GetAllUsers(query)
	if len(users) == 0 {
		utils.ResponseError(ctx, utils.NewError("users not found", utils.ErrCodeNotFound))
		return
	}
	userDTOs := dto.MapUsersToDTOs(users)
	utils.ResponseSuccess(ctx, http.StatusOK, &userDTOs)
}
func (userController *UserController) UpdateUser(ctx *gin.Context) {
	var params dto.UserParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validations.HandleValidationErrors(err))
		return
	}
	var user models.Users
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.ResponseValidator(ctx, validations.HandleValidationErrors(err))
		return
	}
	user, err := userController.service.UpdateUser(params.UUID, user)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	userDTO := dto.MapUserToDTO(user)
	utils.ResponseSuccess(ctx, http.StatusOK, &userDTO)
}
func (userController *UserController) DeleteUser(ctx *gin.Context) {
	var params dto.UserParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validations.HandleValidationErrors(err))
		return
	}
	if err := userController.service.DeleteUser(params.UUID); err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	utils.ResponseSuccess(ctx, http.StatusNoContent, nil)
}
