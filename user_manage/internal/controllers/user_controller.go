package controllers

import (
	"fmt"
	"net/http"
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
		ctx.JSON(http.StatusBadRequest, validations.HandleValidationErrors(err))
		return
	}
	user, err := userController.service.CreateUser(user)
	if err != nil {
		fmt.Println("error:", err)
		utils.ResponseError(ctx, err)
		return
	}
	utils.ResponseSuccess(ctx, http.StatusCreated, user)
}
func (userController *UserController) GetUserByUUID(ctx *gin.Context) {
}
func (userController *UserController) GetAllUsers(ctx *gin.Context) {

}
func (userController *UserController) UpdateUser(ctx *gin.Context) {

}
func (userController *UserController) DeleteUser(ctx *gin.Context) {

}
