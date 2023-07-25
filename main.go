package main

import (
	"fmt"
	"github.com/wayne-campbell/url-shortener/handler"
	"github.com/wayne-campbell/url-shortener/store"
	"github.com/gin-gonic/gin"
	
)

func main(){
	r := gin.Default()
	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Welcome to URL shortener API!",
		})
	})
//API post
r.POST("/create-short-url", func(c *gin.Context){
	handler.CreateShortUrl(c)
})
//API GET
r.GET("/:shortUrl", func(c *gin.Context){
	handler.HandleShortUrlRedirect(c)
})
//initialize store
store.InitializeStore()

err := r.Run(":9808")
if err != nil {
	panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}	
}

