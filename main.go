package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialise the Gin Router
	router := gin.Default()

	// Publish the static files
	router.Static("/public", "./public")

	


	// API to get the Long URL and return the Short URL
	router.GET("/url-shortner", func(ctx *gin.Context) {
		ctx.File("./public/url-shortner.html")
	})

	router.POST("/url-shortner", func(ctx *gin.Context) {
		longUrl := ctx.PostForm("url")
		shortURL := GenerateShortUrl(longUrl)
		ctx.String(http.StatusOK, "Here comes the short URL: "+shortURL)
	})



	// Run the HTTP Server on port 5000
	log.Fatal(router.Run(":5000"))
}
