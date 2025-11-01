package routeV1

import (
	"daniel/gin-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewsHandler struct {
	News string `uri:"news" binding:"required"`
}

func NewNewsHandler() *NewsHandler {
	return &NewsHandler{}
}

func (n *NewsHandler) GetListNews(ctx *gin.Context) {
	var req NewsHandler
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}
	ctx.JSON(200, gin.H{
		"code": "200",
		"news": req.News,
	})
}
