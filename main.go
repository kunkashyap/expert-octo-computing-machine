package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) { //c is gin context
		ctx.JSON(200, gin.H{"message": "hello from gin!"})
	})

	router.GET("/hello/:name", func(ctx *gin.Context) {
		// Path Parameter
		name := ctx.Param("name")

		//Query String
		loud := ctx.DefaultQuery("loud", "false")

		greeting := "hello" + name

		if loud == "true" {
			greeting = strings.ToUpper(greeting)
		}

		ctx.JSON(200, gin.H{"greeting": greeting})
	})

	err := router.Run(":8300")
	if err != nil {
		println("The router encountered an error while running")
	}
}
