package main

import (
	"strconv"

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

	// router.GET("/", func(ctx *gin.Context) { //c is gin context
	// 	ctx.JSON(200, gin.H{"message": "hello from gin!"})
	// })

	// router.GET("/hello/:name", func(ctx *gin.Context) {
	// 	// Path Parameter
	// 	name := ctx.Param("name")

	// 	//Query String
	// 	loud := ctx.DefaultQuery("loud", "false")

	// 	greeting := "hello" + name

	// 	if loud == "true" {
	// 		greeting = strings.ToUpper(greeting)
	// 	}

	// 	ctx.JSON(200, gin.H{"greeting": greeting})
	// })

	//Read all bookmarks

	router.GET("/bookmarks", func(ctx *gin.Context){	})

	//Read one bookmark
	router.GET("/bookmarks/:id", func(ctx *gin.Context) {
		//convert the id from string to int
		id, err := strconv.Atoi(ctx.Param("id")) //convert the id from string to int
		if err != nil { //error handling if the id is not a valid integer
			ctx.JSON(400, gin.H{"error": "Invalid bookmark ID"})
			return
		}
		// Find the bookmark by ID
		for _, bookmark := range bookmarks { //if the bookmark ID matches the requested ID, return it
			if bookmark.ID == id { 
				ctx.JSON(200, bookmark)
				return
			}
		}
		ctx.JSON(404, gin.H{"error": "Bookmark not found"})
	})

	//Update a bookmark
	router.PUT("/bookmarks/:id", func(ctx *gin.Context) {
			id,_ := strconv.Atoi(ctx.Param("id"))
			var updatedBookmark Bookmark
			if err := ctx.ShouldBindJSON(&updatedBookmark); err != nil {
				ctx.JSON(400, gin.H{"error": err.Error()})
				return
			}
			for i, bookmark := range bookmarks {
				if bookmark.ID == id {
					bookmarks[i] = updatedBookmark
					ctx.JSON(200, updatedBookmark)
					return
				}
			}
			ctx.JSON(404, gin.H{"error": "Bookmark not found"})
		})

	//Delete a bookmark
	router.DELETE("/bookmarks/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid bookmark ID"})
			return
		}

		for i, bookmark := range bookmarks {
			if bookmark.ID == id {
				bookmarks = append(bookmarks[:i], bookmarks[i+1:]...)
				ctx.JSON(200, gin.H{"message": "Bookmark deleted"})
				return
			}
		}
		ctx.JSON(404, gin.H{"error": "Bookmark not found"})
	})

	//For creating a new bookmark
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

