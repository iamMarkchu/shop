package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// router  /
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "index2",
	})
}
