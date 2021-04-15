package main

import (
	_ "encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func SayfaninMetniniGetir(konuId int, sayfaNu int) (metin string) {

	url := fmt.Sprintf("http://islamilimleri.com/Ktphn/Kitablar/05/001/Turkce/%02d/%03d.htm", konuId, sayfaNu)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	icerik := doc.Find("#icerik")
	fmt.Println(icerik.Text())

	fmt.Println(doc)
	// Find the review items
	/*doc.Find("#icerik").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		paragraf := s.Find("p").Text()
		fmt.Printf("%s\n")
		fo.WriteString(paragraf)

	})
	doc.Find("#")*/

	return ""
}

func main() {

	SayfaninMetniniGetir(1, 1)
}
