package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

const Turkiye = 2
const RootAdres = "https://ramazan.diyanet.gov.tr/tr-TR/imsakiye/"

func main() {

	tumIlcebilgileri := make([]IlceBilgi, 0, 1000)

	for ilId, il := range Iller() {
		ilceler := IldekiIlceler(Turkiye, ilId)
		for ilceId, ilce := range ilceler {
			ib := IlceBilgi{
				Ulke:     "Türkiye",
				Il:       il,
				Ilce:     ilce,
				UlkeId:   Turkiye,
				IlId:     ilId,
				IlceId:   ilceId,
				Vakitler: IlceImsakiye(Turkiye, ilId, ilceId),
			}
			tumIlcebilgileri = append(tumIlcebilgileri, ib)
			fmt.Println(il, ilce)
		}
	}

	fmt.Println("Eklenen ilce sayısı", len(tumIlcebilgileri))
	if jsonString, err := json.MarshalIndent(tumIlcebilgileri, " ", " "); err != nil {
		log.Fatal(err)
	} else {
		if ioutil.WriteFile("json/imsakiye.json", jsonString, 0666) != nil {
			log.Fatal("imsakiye json yazilamadi")
		}
	}

}
