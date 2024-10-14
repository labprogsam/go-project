package main

import (
	"github.com/labprogsam/go-project/initializers"
	"github.com/labprogsam/go-project/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Bet{})
	initializers.DB.AutoMigrate(&models.UserBet{})
}
