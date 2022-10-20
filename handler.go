package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type shortUrlCreation struct {
	LongUrl string `json:"long-url" binding:"required"`
}

// Create the Short URL on API Request
func createShortUrl(ctx *gin.Context) {
	var urlCreation shortUrlCreation
	if err := ctx.ShouldBindJSON(&urlCreation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if shortUrl, ok := getUrlFromCache(urlCreation.LongUrl); ok {
		log.Printf("From Cache ..%s..", shortUrl)

		ctx.JSON(200, gin.H{
			"message":   "Short Url Created successfully",
			"short-url": shortUrl,
		})
	} else {
		shortUrl = GenerateShortUrl(urlCreation.LongUrl)
		log.Printf("Created New ..%s..", shortUrl)
		setUrlToCache(urlCreation.LongUrl, shortUrl)

		ctx.JSON(200, gin.H{
			"message":   "Short Url Created successfully",
			"short-url": shortUrl,
		})
	}

}
