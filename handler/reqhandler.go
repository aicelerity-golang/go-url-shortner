package handler

import (
	"log"
	"net/http"

	"github.com/aicelerity-golang/go-url-shortner/utils"
	"github.com/gin-gonic/gin"
)

const host = "http://localhost:5000/"

// Define the LongUrl as struct
type Rurldata struct {
	LongUrl string `json:"long-url" binding:"required"`
}

// Create the Short URL on API Request using Fle
func GetShortUrlFile(ctx *gin.Context) {
	var inputURL Rurldata

	if err := ctx.ShouldBindJSON(&inputURL); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if shortUrl, ok := getShortUrlFromStore(inputURL.LongUrl); ok {
		log.Printf("From File ..%s..", shortUrl)

		ctx.JSON(200, gin.H{
			"message":   "Here comes the Short Url...",
			"short-url": host + shortUrl,
		})
	} else {
		shortUrl = utils.GenerateShortUrl(inputURL.LongUrl)
		log.Printf("Created New ..%s..", shortUrl)
		setShortUrlToStore(inputURL.LongUrl, shortUrl)

		ctx.JSON(200, gin.H{
			"message":   "Here comes the Short Url...",
			"short-url": host + shortUrl,
		})
	}

}

func GetLongURLFile(ctx *gin.Context) {
	// var urlCreation shortUrlCreation

	shortURL := ctx.Param("shorturl")

	if longUrl, ok := getLongUrlFromStore(shortURL); ok {
		// log.Printf("From File ..%s..", longUrl)

		ctx.Redirect(302, longUrl)

	} else {

		ctx.JSON(200, gin.H{
			"message": "The Short URL you entered is not in Store",
		})
	}

}

// Create the Short URL on API Request using Cache
func GetShortUrlCache(ctx *gin.Context) {
	var inputURL Rurldata
	if err := ctx.ShouldBindJSON(&inputURL); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if shortUrl, ok := getShortUrlFromCache(inputURL.LongUrl); ok {
		log.Printf("From Cache ..%s..", shortUrl)

		ctx.JSON(200, gin.H{
			"message":   "Short Url Created successfully",
			"short-url": shortUrl,
		})
	} else {
		shortUrl = utils.GenerateShortUrl(inputURL.LongUrl)
		log.Printf("Created New ..%s..", shortUrl)
		setShortUrlToCache(inputURL.LongUrl, shortUrl)

		ctx.JSON(200, gin.H{
			"message":   "Short Url Created successfully",
			"short-url": shortUrl,
		})
	}

}
