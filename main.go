package main

import (
	"log"

	"github.com/aicelerity-golang/go-url-shortner/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	// Initialise the Gin Router
	router := gin.Default()

	// REST API for URL Shortner
	router.POST("/url-shortner", func(ctx *gin.Context) {
		handler.CreateShortUrlFile(ctx)
	})

	router.GET("/:shorturl", func (ctx *gin.Context)  {
		handler.GetLongURLfromFile(ctx)
	} )

	/* 	Basic Web Service to get the Long URL and return the Short URL */
	// Publish the static files
	// router.Static("/public", "./public")

	// router.GET("/url-shortner", func(ctx *gin.Context) {
	// 	ctx.File("./public/url-shortner.html")
	// })

	// router.POST("/url-shortner", func(ctx *gin.Context) {
	// longUrl := ctx.PostForm("url")
	// 	shortURL := createShortUrl(ctx)
	// 	ctx.String(http.StatusOK, "Here comes the short URL: "+shortURL)
	// })

	// Run the HTTP Server on port 5000
	log.Fatal(router.Run(":5000"))
}
