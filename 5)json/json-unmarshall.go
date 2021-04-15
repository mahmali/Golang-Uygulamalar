package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	type tip struct {
		İsim    string `json:"İsim"`
		Soyisim string `json:"Soyisim"`
		Yas     int    `json:"Yas"`
	}
	var veri tip
	veri.Yas = 3
	veri.Soyisim = "omer"
	veri.İsim = "tuncer"
	veri1 := tip{
		İsim:    "omer",
		Soyisim: "faruk",
		Yas:     8,
	}
	olusan, _ := json.Marshal(&veri)
	olusan2, _ := json.Marshal(&veri1)
	fmt.Println(string(olusan))
	fmt.Println(string(olusan2))

	var dat []tip
	donen, err := ioutil.ReadFile("C:\\Users\\Muhammed Ali AKKAYA\\Desktop\\golang2\\json\\data.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	json.Unmarshal(donen, &dat)
	fmt.Println("cenabet ali")
}
