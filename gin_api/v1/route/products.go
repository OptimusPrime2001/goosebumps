package routeV1

import "github.com/gin-gonic/gin"

type ProductRoute struct {
}

func NewProductRoute() *ProductRoute {
	return &ProductRoute{}
}

func (u *ProductRoute) GetListProduct(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code": "200",
		"data": "List product version 1",
	})
}
func (u *ProductRoute) GetProductByID(ctx *gin.Context) {
	product_id := ctx.Params.ByName("product_id")
	ctx.JSON(200, gin.H{
		"code":       "200",
		"product_id": product_id,
	})
}
func (u *ProductRoute) DeleteProduct(ctx *gin.Context) {
	product_id := ctx.Params.ByName("product_id")
	ctx.JSON(200, gin.H{
		"code":       "200",
		"product_id": product_id,
	})
}
func (u *ProductRoute) UpdateProduct(ctx *gin.Context) {
	product_id := ctx.Params.ByName("product_id")
	ctx.JSON(200, gin.H{
		"code":       "200",
		"product_id": product_id,
	})
}
