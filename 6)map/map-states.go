package main

import "fmt"

func main() {
	states := make(map[string]string)
	fmt.Println(states)
	states["ist"] = "istanbul"
	states["ank"] = "ankara"
	states["mar"] = "kahraman maraş"
	fmt.Println(states)

	//verilen map listesinden  ank adına sahip veriyi else etme işlemi
	ankara := states["ank"]
	fmt.Println(ankara)

	//veri silinme işlemi
	delete(states, "ank")
	fmt.Println(states)

	//ekleme işlemi
	states["elz"] = "elazıg"

	for key, value := range states { //burada i ye map in key degeri atanır. normal dilerdeki gibi 0 1 2 gibi sayısal sıralama işlemi olmaz
		fmt.Printf("map in %v. elemanı %v \n", key, value)
	}

	//state map indeki keyler başka bir listeye atma işlemi yapılacak
	keys := make([]string, len(states)) //states in boyutunda bir liste tanımlama işlemi yapıldı
	i := 0
	for key := range states {
		keys[i] = key
		i++
	}
	fmt.Println("keys listesi elemanları", keys)

	//key listesindeki index verilere göre states map indeki elemanları yazdırma işlemi yapma
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
