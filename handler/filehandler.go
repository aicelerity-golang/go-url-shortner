package handler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// Define the Urldta as struct
type Nurldata struct {
	NlongUrl  string
	NshortUrl string
}

// Get the Short Url from Store
func getShortUrlFromStore(longUrl string) (string, bool) {

	ok := false
	shortUrl := getUrlFromFile(longUrl)
	if shortUrl != "" {
		ok = true
	}
	return shortUrl, ok
}

func getLongUrlFromStore(shortURL string) (string, bool) {

	ok := false
	longUrl := getUrlFromFile(shortURL)
	if longUrl != "" {
		ok = true
	}
	return longUrl, ok
}

// Update the  Short Url in Store
func setShortUrlToStore(longUrl, shortUrl string) {

	urls := urlFileReader()
	urlFileWriter(urls, longUrl, shortUrl)
	log.Printf("The Short URL ..%s.. is stored in File", shortUrl)
}

// Reads the file and return the Short URL

func getUrlFromFile(url string) string {

	f, err := os.Open("urldata.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ""
	}
	defer f.Close()

	buf := make([]byte, 0, 1024)
	r := bufio.NewReader(f)
	var foundUrl string
	for {
		n, err := io.ReadFull(r, buf[:cap(buf)])
		buf = buf[:n]
		if err != nil {
			if err == io.EOF {
				break
			}
			if err != io.ErrUnexpectedEOF {
				fmt.Fprintln(os.Stderr, err)
				break
			}
		}

		var urls []Nurldata
		// Unmarshall the data
		err = json.Unmarshal(buf, &urls)
		if err != nil {
			log.Fatal(err)
		}
		foundUrl = searchURL(urls, url)
	}
	return foundUrl
}

func searchURL(urls []Nurldata, url string) string {
	var foundUrl string
	if len(url) <= 10 {
		for i := 0; i < len(urls); i++ {
			if urls[i].NshortUrl == url {
				foundUrl = urls[i].NlongUrl
			}
		}
	} else {
		for i := 0; i < len(urls); i++ {
			if urls[i].NlongUrl == url {
				foundUrl = urls[i].NshortUrl
			}
		}
	}
	return foundUrl
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
