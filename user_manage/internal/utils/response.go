package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseError(ctx *gin.Context, err error) {
    if appErr, ok := err.(*AppError); ok {
        statusCode := httpStatusFromErrorCode(appErr.Code)
        response := gin.H{
            "code": appErr.Code,
            "msg":  appErr.Message,
        }
        if appErr.Err != nil {
            response["detail"] = appErr.Err.Error()
        }
        ctx.AbortWithStatusJSON(statusCode, response)
        return
    }
    ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
        "code": ErrCodeInternal,
        "msg":  err.Error(),
    })
}

func ResponseSuccess(ctx *gin.Context, status int, data any) {
	ctx.JSON(status, gin.H{
		"code": status,
		"msg":  "success",
		"data": data,
	})
}
