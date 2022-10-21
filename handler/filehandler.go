package handler

import (
	"encoding/json"
	"log"
	"os"
)

// Define the Urldta as struct
type Nurldata struct {
	NlongUrl  string
	NshortUrl string
}

// Get the Short Url from Store
func GetUrlFromStore(longUrl string) (string, bool) {

	ok := false
	shortUrl := GetShortUrlFromFile(longUrl)
	// log.Printf("The Short URL is : %s", shortUrl)
	if shortUrl != "" {
		ok = true
	}
	return shortUrl, ok
}

// Update the  Short Url in Store
func setUrlToStore(longUrl, shortUrl string) {

	urls := urlFileReader()
	urlFileWriter(urls, longUrl, shortUrl)
	log.Printf("The Short URL ..%s.. is stored in File", shortUrl)
}

// Reads the file and return the Short URL
func GetShortUrlFromFile(longUrl string) string {

	content, err := os.ReadFile("urldata.txt")
	if err != nil {
		log.Fatal(err)
	}
	var urls []Nurldata

	// Unmarshall the data
	err = json.Unmarshal(content, &urls)
	if err != nil {
		log.Fatal(err)
	}

	var shortUrl string
	for i := 0; i < len(urls); i++ {
		if urls[i].NlongUrl == longUrl {
			// log.Println(urls[i].NshortUrl)
			shortUrl = urls[i].NshortUrl
		}
	}
	return shortUrl
}

// Reads the file and return all urls
func urlFileReader() []Nurldata {
	content, err := os.ReadFile("urldata.txt")
	if err != nil {
		log.Fatal(err)
	}
	var urls []Nurldata

	// Unmarshall the data
	err = json.Unmarshal(content, &urls)
	if err != nil {
		log.Fatal(err)
	}
	return urls
}

// Write all urls to File
func urlFileWriter(urls []Nurldata, longUrl, shortUrl string) {

	urls = append(urls, Nurldata{NlongUrl: longUrl, NshortUrl: shortUrl})

	// Marshal the data
	content, err := json.Marshal(urls)
	if err != nil {
		log.Println(err)
	}

	// Write to file
	err = os.WriteFile("urldata.txt", content, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
