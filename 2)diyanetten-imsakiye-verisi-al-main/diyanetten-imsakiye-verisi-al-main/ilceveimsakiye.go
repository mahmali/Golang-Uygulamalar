package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func IldekiIlceler(ulkeId, ilId int) map[int]string {

	data := url.Values{
		"ulkeId": {strconv.Itoa(ulkeId)},
		"ilId":   {strconv.Itoa(ilId)},
		"ilceId": {},
	}

	resp, err := http.PostForm(RootAdres, data)

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	ret := make(map[int]string)
	doc.Find("#ilceId option").Each(func(i int, s *goquery.Selection) {
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

func IlceImsakiye(ulkeId, ilId, ilceId int) []Imsakiye {

	ret := make([]Imsakiye, 0, 30)
	data := url.Values{
		"ulkeId": {strconv.Itoa(ulkeId)},
		"ilId":   {strconv.Itoa(ilId)},
		"ilceId": {strconv.Itoa(ilceId)},
	}

	resp, err := http.PostForm(RootAdres, data)

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#table tr").Each(func(i int, tr *goquery.Selection) {
		if i < 2 || i > 32 {
			return
		}
		imsakiye := Imsakiye{}
		tr.Find("td").Each(func(ti int, td *goquery.Selection) {
			txt := strings.TrimSpace(td.Text())
			switch ti {
			case 0:
				imsakiye.HicriTarih = txt
			case 1:
				imsakiye.MiladiTarih = txt
			case 2:
				imsakiye.Imsak = txt
			case 3:
				imsakiye.Gunes = txt
			case 4:
				imsakiye.Ogle = txt
			case 5:
				imsakiye.Ikindi = txt
			case 6:
				imsakiye.Aksam = txt
			case 7:
				imsakiye.Yatsi = txt
			}
		})
		if imsakiye.HicriTarih != "KADİR GECESİ" {
			ret = append(ret, imsakiye)
		}

	})

	return ret
}
