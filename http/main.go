package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		return "", fmt.Errorf("Couldn't find a content-type")
	}
	return contentType, nil
}

func main() {
	contentType, err := contentType("https://golang.org")

	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		fmt.Println(contentType)
	}
}
