package handler

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlsFromStore(t *testing.T) {

	log.Println("testing case 1 for getShortUrlFromFile")
	initialUrl_1 := "https://blogs.blackberry.com/en/2022/10/bianlian-ransomware-encrypts-files-in-the-blink-of-an-eye"
	shortUrl_1 := GetShortUrlFromFile(initialUrl_1)

	assert.Equal(t, shortUrl_1, "RX9wbW9WtQ")

	log.Println("testing case 2 for getShortUrlFromFile")
	initialUrl_2 := "https://www.infosecurity-magazine.com/news/golang-malware-image-james-webb/"
	shortUrl_2 := GetShortUrlFromFile(initialUrl_2)

	assert.Equal(t, shortUrl_2, "")

	log.Println("testing case 3 for getUrlFromStore")
	initialUrl_3 := "https://blogs.blackberry.com/en/2022/10/bianlian-ransomware-encrypts-files-in-the-blink-of-an-eye"
	shortUrl_3, status_3 := GetUrlFromStore(initialUrl_3)

	assert.Equal(t, shortUrl_3, "RX9wbW9WtQ")
	assert.Equal(t, status_3, true)

	log.Println("testing case 4 for getUrlFromStore")
	initialUrl_4 := "https://www.infosecurity-magazine.com/news/golang-malware-image-james-webb/"
	shortUrl_4, status_4 := GetUrlFromStore(initialUrl_4)

	assert.Equal(t, shortUrl_4, "")
	assert.Equal(t, status_4, false)

}
