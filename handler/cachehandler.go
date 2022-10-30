package handler

import "log"

var cURLData = make(map[string]string)

// Initialise the Map - cURLData with some data
func init() {

	var bigUrl = "https://blogs.blackberry.com/en/2022/10/bianlian-ransomware-encrypts-files-in-the-blink-of-an-eye"
	var shortUrl = "RX9wbW9WtQ"

	cURLData[bigUrl] = shortUrl

}

// get the short URL from Cache if availabe.
func getShortUrlFromCache(longUrl string) (string, bool) {

	shortUrl, ok := cURLData[longUrl]
	// log.Println("The Short URL is : ", shortUrl)
	return shortUrl, ok
}

// Set the short URL to Cache
func setShortUrlToCache(longUrl, shortUrl string) {

	cURLData[longUrl] = shortUrl
	log.Printf("The Short URL ..%s.. is stored in Cache", shortUrl)
}
