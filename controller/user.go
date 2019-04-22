package controller

import (
	"log"

	"huochain/gomicro-demo/api/pb"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var (
	userService pb.UserService
)

// 用户列表
func (c *UserController) Index(ctx *gin.Context) {
	log.Print("request usercontroller index")
	ctx.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

// 获取用户信息
func (c *UserController) Info(ctx *gin.Context) {
	log.Print("request user info from userid " + ctx.Param("userId"))
	user, err := userService.Info(ctx.TODO(), &pb.User() {
		Id: 123
	})

	if err != nil {
		log.Fatal("call userService.Info error: " + err)

		ctx.JSON(500, map[string]interface{}{
			"code": 500,
			"message": "internal server error",
		})
	}

	ctx.JSON(200, map[string]interface{}{
		"code": 100,
		"message": "",
		"data": user
	})
}
