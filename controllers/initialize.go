package controllers

import (
	"github.com/gin-gonic/gin"
)

func InitializeControllers(r *gin.Engine) {
	r.GET("/", IndexGet)
	r.POST("/", IndexPost)
	r.GET("/test", TestGet)
}
