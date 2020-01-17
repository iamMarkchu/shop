package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/shop/services"
	"net/http"
)

// @router /user/register [post]
// @param  userName string required
// @param  password string required
// @param  rePassword string required
// @param  captcha  string
func Register(c *gin.Context) {
	userName := c.PostForm("userName")
	passWord := c.PostForm("password")
	rePassword := c.PostForm("rePassword")
	userService := services.NewUserService()
	result := userService.Register(userName, passWord, rePassword)
	c.JSON(http.StatusOK, gin.H{
		"api":    "register",
		"result": result,
	})
}

// router /user/login [post]
// @param  userName string required
// @param  password string required
// @param  captcha  string
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"api": "login",
	})
}
