package config

import (
	"huochain/gomicro-demo/api/controller"

	"github.com/gin-gonic/gin"
)

func BuildRouter() *gin.Engine {
	userController := new(controller.UserController)

	router := gin.Default()
	router.GET("/user", userController.Index)
	router.GET("/user/:userId", userController.Info)

	return router
}
