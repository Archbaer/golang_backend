package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"backend/middlewares"
	"backend/routes"
)

func main() {
	fmt.Println("Starting...")

	r := gin.Default()

	r.Use(middlewares.Logging())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, You are at root!",
		})
	})

	r.GET("/dominican", func(c *gin.Context) {
		fmt.Println(" I no black, I dominican papi!")
	}, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "You dominican papi!",
		})
	})

	routes.UserRoute(r)

	r.Run(":3333")
}
