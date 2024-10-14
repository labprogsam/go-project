package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labprogsam/go-project/initializers"
	"github.com/labprogsam/go-project/models"
)

func UsersCreate(c *gin.Context) {
	// get data off req body
	var body struct {
		Name      string
		Password  string
		Username  string
		IsAdmin   bool
		Pontuacao float64
	}

	c.Bind(&body)

	user := models.User{
		Name:      body.Name,
		Password:  body.Password,
		Username:  body.Username,
		IsAdmin:   body.IsAdmin,
		Pontuacao: body.Pontuacao,
	}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
