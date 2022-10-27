package handler

import (
	"log"
	"net/http"

	"github.com/aicelerity-golang/go-url-shortner/utils"
	"github.com/gin-gonic/gin"
)

const host = "http://localhost:5000/"

// Define the LongUrl as struct
type shortUrlCreation struct {
	LongUrl string `json:"long-url" binding:"required"`
}

// Create the Short URL on API Request using Fle
func CreateShortUrlFile(ctx *gin.Context) {
	var urlCreation shortUrlCreation

	if err := ctx.ShouldBindJSON(&urlCreation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if shortUrl, ok := GetUrlFromStore(urlCreation.LongUrl); ok {
		log.Printf("From File ..%s..", shortUrl)

		ctx.JSON(200, gin.H{
			"message":   "Here comes the Short Url...",
			"short-url": host + shortUrl,
		})
	} else {
		shortUrl = utils.GenerateShortUrl(urlCreation.LongUrl)
		log.Printf("Created New ..%s..", shortUrl)
		setUrlToStore(urlCreation.LongUrl, shortUrl)

		ctx.JSON(200, gin.H{
			"message":   "Here comes the Short Url...",
			"short-url": host + shortUrl,
		})
	}

}

func GetLongURLfromFile(ctx *gin.Context) {
	// var urlCreation shortUrlCreation

	shortURL := ctx.Param("shorturl")

	if longUrl, ok := GetLongUrlFromStore(shortURL); ok {
		// log.Printf("From File ..%s..", longUrl)

		ctx.Redirect(302, longUrl)

	} else {

		ctx.JSON(200, gin.H{
			"message": "The Short URL you entered is not in Store",
		})
	}

}

// Create the Short URL on API Request using Cache
// func CreateShortUrlCache(ctx *gin.Context) {
// 	var urlCreation shortUrlCreation
// 	if err := ctx.ShouldBindJSON(&urlCreation); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if shortUrl, ok := getUrlFromCache(urlCreation.LongUrl); ok {
// 		log.Printf("From Cache ..%s..", shortUrl)

// 		ctx.JSON(200, gin.H{
// 			"message":   "Short Url Created successfully",
// 			"short-url": shortUrl,
// 		})
// 	} else {
// 		shortUrl = GenerateShortUrl(urlCreation.LongUrl)
// 		log.Printf("Created New ..%s..", shortUrl)
// 		setUrlToCache(urlCreation.LongUrl, shortUrl)

// 		ctx.JSON(200, gin.H{
// 			"message":   "Short Url Created successfully",
// 			"short-url": shortUrl,
// 		})
// 	}

// }
