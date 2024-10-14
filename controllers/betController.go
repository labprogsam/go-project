package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labprogsam/go-project/initializers"
	"github.com/labprogsam/go-project/models"
)

func BetsCreate(c *gin.Context) {
	// get data off req body
	var body struct {
		Team1 string
		Team2 string
		Odd1  float64
		Odd2  float64
	}

	c.Bind(&body)

	bet := models.Bet{
		Team1:  body.Team1,
		Team2:  body.Team2,
		Odd1:   body.Odd1,
		Odd2:   body.Odd2,
		Result: 0,
		IsOpen: true,
	}
	result := initializers.DB.Create(&bet)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"bet": bet,
	})
}
