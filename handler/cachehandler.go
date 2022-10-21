package handler

import "log"

var urlData = make(map[string]string)

// Initialise the Map - urlData with some data
func init() {

	var bigUrl = "https://blogs.blackberry.com/en/2022/10/bianlian-ransomware-encrypts-files-in-the-blink-of-an-eye"
	var shortUrl = "RX9wbW9WtQ"

	urlData[bigUrl] = shortUrl

}

// get the short URL from Cache if availabe.
func getUrlFromCache(longUrl string) (string, bool) {

	shortUrl, ok := urlData[longUrl]
	// log.Println("The Short URL is : ", shortUrl)
	return shortUrl, ok
}

// Set the short URL to Cache
func setUrlToCache(longUrl, shortUrl string) {

	urlData[longUrl] = shortUrl
	log.Printf("The Short URL ..%s.. is stored in Cache", shortUrl)
}
