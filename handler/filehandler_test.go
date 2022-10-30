package handler

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlsFromStore(t *testing.T) {

	log.Println("Testing case 1 for getUrlFromFile")
	longUrl_1 := "https://blogs.blackberry.com/en/2022/10/bianlian-ransomware-encrypts-files-in-the-blink-of-an-eye"
	shortUrl_1 := getUrlFromFile(longUrl_1)

	assert.Equal(t, shortUrl_1, "RX9wbW9WtQ")

	log.Println("Testing case 2 for getUrlFromFile")
	longUrl_2 := "https://www.infosecurity-magazine.com/news/golang-malware-image-james-webb/"
	shortUrl_2 := getUrlFromFile(longUrl_2)

	assert.Equal(t, shortUrl_2, "")

	log.Println("Testing case 3 for getShortUrlFromStore")
	longUrl_3 := "https://blogs.blackberry.com/en/2022/10/bianlian-ransomware-encrypts-files-in-the-blink-of-an-eye"
	shortUrl_3, status_3 := getShortUrlFromStore(longUrl_3)

	assert.Equal(t, shortUrl_3, "RX9wbW9WtQ")
	assert.Equal(t, status_3, true)

	log.Println("Testing case 4 for getShortUrlFromStore")
	longUrl_4 := "https://www.infosecurity-magazine.com/news/golang-malware-image-james-webb/"
	shortUrl_4, status_4 := getShortUrlFromStore(longUrl_4)

	assert.Equal(t, shortUrl_4, "")
	assert.Equal(t, status_4, false)

	log.Println("Testing case 5 for getUrlFromFile")
	shortUrl_5 := "RX9wbW9WtQ"
	longUrl_5 := getUrlFromFile(shortUrl_5)

	assert.Equal(t, longUrl_5, "https://blogs.blackberry.com/en/2022/10/bianlian-ransomware-encrypts-files-in-the-blink-of-an-eye")

	log.Println("Testing case 6 for getUrlFromFile")
	shortUrl_6 := "JH9wbW9WdR"
	longUrl_6 := getUrlFromFile(shortUrl_6)

	assert.Equal(t, longUrl_6, "")

}
