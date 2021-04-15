package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Make request
	response, err := http.Get("http://islamilimleri.com/Ktphn/Kitablar/05/001/Turkce/01/001.htm")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create output file
	outFile, err := os.Create("output.html")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// Copy data from HTTP response to file
	_, err = io.Copy(outFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
}
