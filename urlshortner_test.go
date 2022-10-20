package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortUrls(t *testing.T) {
	initialUrl_1 := "https://blogs.blackberry.com/en/2022/10/bianlian-ransomware-encrypts-files-in-the-blink-of-an-eye"
	shortUrl_1 := GenerateShortUrl(initialUrl_1)

	initialUrl_2 := "https://www.infosecurity-magazine.com/news/golang-malware-image-james-webb/"
	shortUrl_2 := GenerateShortUrl(initialUrl_2)

	assert.Equal(t, shortUrl_1, "RX9wbW9WtQ")
	assert.Equal(t, shortUrl_2, "4P-ZgXZQPF")
}
