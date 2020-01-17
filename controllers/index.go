package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// router  /
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"page": "index",
	})
}
