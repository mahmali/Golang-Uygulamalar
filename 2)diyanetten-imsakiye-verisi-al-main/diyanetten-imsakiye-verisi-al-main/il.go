package main

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func Iller() map[int]string {

	var iller map[int]string
	if _, err := os.Stat("json/iller.json"); os.IsNotExist(err) {
		iller = webdenIlleriGetir()
		if jsonString, err := json.MarshalIndent(iller, " ", " "); err != nil {
			log.Fatal(err)
		} else {
			if ioutil.WriteFile("json/iller.json", jsonString, 0666) != nil {
				log.Fatal("iller json yazilamadi")
			}
		}

	} else {
		f, _ := os.Open("json/iller.json")
		if json.NewDecoder(f).Decode(&iller) != nil {
			log.Fatal("iller jsondan alinamadı")
		}
	}
	return iller
}

func webdenIlleriGetir() map[int]string {
	res, err := http.Get(RootAdres)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	ret := make(map[int]string)
	doc.Find("#ilId option").Each(func(i int, s *goquery.Selection) {
		val, _ := s.Attr("value")
		if val == "" || s.Text() == "" {
			return
		}
		id, _ := strconv.Atoi(val)
		if id > 0 {
			ret[id] = s.Text()
		}
	})
	return ret
}
