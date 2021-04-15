package main

import (
	_ "encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strings"
)

func ExampleScrape(deger int) {
	// Request the HTML page.
	var dosya string = fmt.Sprintf("dosya%v.txt", deger)
	fo, err := os.Create(dosya)
	if err != nil {
		panic(err)
	}
	var url string
	if deger < 10 {
		url = fmt.Sprintf("http://islamilimleri.com/Ktphn/Kitablar/05/001/Turkce/0%v/000.htm", deger)
	} else {
		url = fmt.Sprintf("http://islamilimleri.com/Ktphn/Kitablar/05/001/Turkce/%v/000.htm", deger)
	}
	//url := fmt.Sprintf("http://islamilimleri.com/Ktphn/Kitablar/05/001/Turkce/0%v/002.htm", deger)
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	req.Header.Add("Accept-Charset", "text/html; charset=windows-1254;")
	req.Header.Add("Content-Type", "text/html; charset=windows-1254;")
	resp, err := http.DefaultClient.Do(req)

	/*	res, err := http.Get(url)
		//res, err := http.Get("http://islamilimleri.com/Ktphn/Kitablar/05/001/Turkce/02/002.htm")
		if err != nil {
			log.Fatal(err)
		}*/
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(doc)
	// Find the review items
	doc.Find("#icerik").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		paragraf := s.Find("p").Text()
		fmt.Printf("%s\n")
		fo.WriteString(paragraf)
	})

}

func main() {
	var i int
	for i = 1; i < 5; i++ {
		ExampleScrape(i)
	}
	//ExampleScrape(3)
}
