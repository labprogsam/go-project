package main

import (
	"github.com/gin-gonic/gin"
	"github.com/labprogsam/go-project/controllers"
	"github.com/labprogsam/go-project/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/bets", controllers.BetsCreate)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
