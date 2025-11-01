package routeV1

import (
	"daniel/gin-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserRoute struct {
}
type GetUserByIDRequest struct {
	UserID int `uri:"user_id" binding:"gt=0"`
}
type UserFamily struct {
	ParentID int `json:"parent_id" binding:"required,gt=0"`
	ChildID  int `json:"child_id" binding:"required,gt=0"`
}
type AddUserPayload struct {
	Name   string     `json:"name" binding:"required,min=3,max=20"`
	Email  string     `json:"email" binding:"required,email"`
	Age    int        `json:"age" binding:"required,gt=0,lt=100"`
	Family UserFamily `json:"family" binding:"required"`
}

func NewUserRoute() *UserRoute {
	return &UserRoute{}
}

func (u *UserRoute) GetListUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code": "200",
		"data": "List user version 1",
	})
}
func (u *UserRoute) GetUserByID(ctx *gin.Context) {
	var req GetUserByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}
	ctx.JSON(200, gin.H{
		"code":    "200",
		"user_id": req.UserID,
	})
}
func (u *UserRoute) DeleteUser(ctx *gin.Context) {
	user_id := ctx.Params.ByName("user_id")
	id, err := strconv.Atoi(user_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "User_id must be a number",
		})
	}
	if id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{})

	}
	ctx.JSON(200, gin.H{
		"code":    "200",
		"user_id": user_id,
	})
}

func (u *UserRoute) UpdateUser(ctx *gin.Context) {
	var payload AddUserPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}
	ctx.JSON(200, gin.H{
		"code": "200",
		"data": payload,
	})
}
