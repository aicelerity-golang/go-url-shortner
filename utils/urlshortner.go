package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

// Hash the input string to get a unique value
func sha256Hash(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

// Use Base64 to Encode the Hashed data
func urlEncoding(data string) string {
	byteData := []byte(data)
	encoded := base64.RawURLEncoding.EncodeToString(byteData)
	return string(encoded)
}

// Generate a Shortened URL for the given input URL
func GenerateShortUrl(inputUrl string) string {
	urlHashBytes := sha256Hash(inputUrl)
	shortURL := urlEncoding(string(urlHashBytes))
	// log.Println(shortURL[:10])
	return shortURL[:10]
}
