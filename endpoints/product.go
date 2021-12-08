package endpoints

import (
	controllers "warehouse/controllers/product"

	"github.com/gin-gonic/gin"
)

func ActivateCategory(router *gin.Engine) {
	categoryRepository := controllers.NewCategory()
	router.POST("/categories", categoryRepository.CreateCategory)
	router.PUT("/categories/:id", categoryRepository.UpdateCategory)
	router.DELETE("/categories/:id", categoryRepository.DeleteCategory)
	router.GET("/categories", categoryRepository.GetCategories)
	router.GET("/categories/:id", categoryRepository.GetCategory)
}

func ActivateDiscount(router *gin.Engine) {
	discountRepository := controllers.NewDiscount()
	router.POST("/discounts", discountRepository.CreateDiscount)
	router.PUT("/discounts/:id", discountRepository.UpdateDiscount)
	router.DELETE("/discounts/:id", discountRepository.DeleteDiscount)
	router.GET("/discounts", discountRepository.GetDiscounts)
	router.GET("/discounts/:id", discountRepository.GetDiscount)
}

func ActivateProduct(router *gin.Engine) {
	productRepository := controllers.NewProduct()
	router.POST("/products", productRepository.CreateProduct)
	router.PUT("/products/:id", productRepository.UpdateProduct)
	router.DELETE("/products/:id", productRepository.DeleteProduct)
	router.GET("/products", productRepository.GetProducts)
	router.GET("/products/:id", productRepository.GetProduct)
}
