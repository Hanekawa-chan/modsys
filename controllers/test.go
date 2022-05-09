package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestGet(c *gin.Context) {
	c.HTML(http.StatusOK,
		// Use the index.html template
		"test.html",
		// Pass the data that the page uses (in this case, 'title')
		nil)
}
