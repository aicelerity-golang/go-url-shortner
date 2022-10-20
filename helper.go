package main

import "log"

var urlData = make(map[string]string)

// Initialise the Map - urlData with some data
func init() {

	var bigUrl = "https://blogs.blackberry.com/en/2022/10/bianlian-ransomware-encrypts-files-in-the-blink-of-an-eye"
	var shortUrl = "RX9wbW9WtQ"

	urlData[bigUrl] = shortUrl

	// for key, value := range urlData {
	// 	log.Println("Key:", key, "Value:", value)
	// }
}

// get the short URL from Cache if availabe.
func getUrlFromCache(longUrl string) {

	shortUrl := urlData[longUrl]
	log.Println("The Short URL is : ", shortUrl)
}

// Set the short URL to Cache
func setUrlToCache(longUrl, shortUrl string) {

	urlData[longUrl] = shortUrl
	log.Println("The Short URL is stored: ", shortUrl)
}
