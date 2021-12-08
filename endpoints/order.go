package endpoints

import (
	controllers "warehouse/controllers"

	"github.com/gin-gonic/gin"
)

func ActivateOrder(router *gin.Engine) {
	orderRepository := controllers.NewOrder()
	router.POST("/orders/:change", orderRepository.ChangeAmountOfOrder)
	router.GET("/orders/:id/user", orderRepository.GetOrders)
}
