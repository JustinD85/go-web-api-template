package main

import (
	"github.com/gin-gonic/gin"
	"github.com/justind85/go-web-api-template/initializers"
)

func init() { //special function, it will run as initializer automagically
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default() //router

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
