package main

import (
	"daniel/gin-api/middleware"
	routeV1 "daniel/gin-api/v1/route"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env if present
	if err := godotenv.Load(".env"); err != nil {
		log.Println("failed to load .env:", err)
	}
	server := gin.Default()
	go middleware.CleanupRateLimiters()
	server.Use(
		middleware.ApiKeyMiddleware(),
		middleware.RateLimitMiddleware(),
	)
	server.Use(middleware.SimpleMiddleware())
	v1 := server.Group("/api/v1")
	{
		user := v1.Group("/user")
		{
			userRoute := routeV1.NewUserRoute()
			user.GET("", userRoute.GetListUser)
			user.GET("/:user_id", userRoute.GetUserByID)
			user.DELETE("/:user_id", userRoute.DeleteUser)
			user.PUT("/:user_id", userRoute.UpdateUser)
		}

		product := v1.Group("/product")
		{
			productRoute := routeV1.NewProductRoute()
			product.GET("", productRoute.GetListProduct)
			product.GET("/:product_id", productRoute.GetProductByID)
			product.DELETE("/:product_id", productRoute.DeleteProduct)
			product.PUT("/:product_id", productRoute.UpdateProduct)

		}

		category := v1.Group("/categories")
		{
			categoryRoute := routeV1.NewCategoryRoute()
			category.GET("/:category", middleware.RateLimitMiddleware(), categoryRoute.GetListCategory)
			// category.GET("/:category_id", categoryRoute.GetCategoryByID)
			// category.DELETE("/:category_id", categoryRoute.DeleteCategory)
			// category.PUT("/:category_id", categoryRoute.UpdateCategory)
			category.POST("", categoryRoute.PostCategory)
		}

		news := v1.Group("/news")
		{
			newsHandler := routeV1.NewNewsHandler()
			news.GET("", newsHandler.GetListNews)
		}
	}
	server.Run("127.0.0.1:8080")
}
