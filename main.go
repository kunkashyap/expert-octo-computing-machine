package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/",func(c *gin.Context) { //c is gin context
		c.JSON(200, gin.H{"message": "hello from gin!"})
	})

	err := router.Run(":8080")
	if err != nil {
		println("The router encountered an error while running")
	}
}
