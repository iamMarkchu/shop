package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/shop/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/", controllers.Index)
	// router user group
	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", controllers.Register)
		userRouter.POST("/login", controllers.Login)
	}

	if err := r.Run(); err != nil {
		fmt.Println("启动出现错误:", err.Error())
	}
}
