package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)



type Bookmark struct {
	ID int `json:"id"`
	Title string `json:"title" binding:"required"`
	URL string `json:"url" binding:"required,url"` //checks non empty feild and valid url
}

var bookmarks = []Bookmark{}
var nextID = 1


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

	router.POST("/bookmarks", func(ctx *gin.Context) {
		var newBookmark Bookmark


		//Json Binding
		if err := ctx.ShouldBindJSON(&newBookmark); err != nil{
				ctx.JSON(400, gin.H{"error": err.Error()})
				return
		}
		newBookmark.ID = nextID
		nextID++ // increment nextID for the next bookmark

		bookmarks = append(bookmarks, newBookmark)
		ctx.JSON(201,newBookmark ) // return the newly created bookmark with a 201 status code
	})

	err := router.Run(":8300") // listen and serve on port 8300
	if err != nil {
		println("The router encountered an error while running")
	}
}

