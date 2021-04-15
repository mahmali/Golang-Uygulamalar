package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type tip struct {
	İsim    string `json:"İsim"`
	Soyisim string `json:"Soyisim"`
	Yas     int    `json:"Yas"`
}

func main() {

	insanlar := make([]tip, 0)

	insan := tip{İsim: "muhammed", Soyisim: "akkaya", Yas: 20}

	for i := 0; i < 10; i++ {
		insan := tip{İsim: "muhammed", Soyisim: "akkaya", Yas: i}
		insanlar = append(insanlar, insan)
	}

	insanlar = append(insanlar, insan)

	fmt.Println(insan)

	data, _ := json.MarshalIndent(insanlar, " ", " ")

	ioutil.WriteFile("data.json", data, 0644)
}
