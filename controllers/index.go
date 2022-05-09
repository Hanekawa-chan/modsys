package controllers

import (
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

func IndexGet(c *gin.Context) {
	c.HTML(http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses (in this case, 'title')
		nil)
}

func IndexPost(c *gin.Context) {
	name := c.PostForm("name")
	surname := c.PostForm("surname")
	number, err := strconv.Atoi(c.PostForm("number"))
	if err != nil {
		log.Fatal().Err(err).Msg("Couldn't convert number to int")
	}
	user := models.NewUser(name, surname, number)
	log.Info().Msg(user.ToString())
	c.Redirect(http.StatusOK, "localhost:8080/test")
}
