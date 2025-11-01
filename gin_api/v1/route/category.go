package routeV1

import (
	"daniel/gin-api/utils"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type CategoryRoute struct {
}

type GetCategoryParam struct {
	Category string `uri:"category" binding:"oneof=php python golang"`
}
type PostCategoryParam struct {
	Name   string `form:"name" binding:"required"`
	Status string `form:"status" binding:"required"`
}

func NewCategoryRoute() *CategoryRoute {
	return &CategoryRoute{}
}
func (c *CategoryRoute) GetListCategory(ctx *gin.Context) {
	var req GetCategoryParam
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}
	ctx.JSON(200, gin.H{
		"code":     "200",
		"category": req.Category,
	})
}
func (c *CategoryRoute) PostCategory(ctx *gin.Context) {
	var req PostCategoryParam
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}
	image, nil := ctx.FormFile("image")
	if err1 := utils.CheckFileName(image); err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg":  err1.Error(),
		})
		return
	}
	if err1 := utils.CheckFileSize(image, 1<<20); err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg":  err1.Error(),
		})
		return
	}
	// os.ModePerm is 0777 (octal)
	// Meaning you have permission to read, write, and execute the folder
	err1 := os.MkdirAll("./upload", os.ModePerm)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": "500",
			"msg":  "create upload folder failed",
		})
		return
	}
	des := fmt.Sprintf("./upload/%s", filepath.Base(image.Filename))
	err1 = ctx.SaveUploadedFile(image, des)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": "500",
			"msg":  "save image failed",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":     "200",
		"category": req.Name,
		"status":   req.Status,
		"image":    des,
	})
}
