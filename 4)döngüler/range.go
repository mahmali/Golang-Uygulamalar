package main

import "fmt"

func main() {
	/*
		var pow = [] int {1,2,4,8,16,32,64,128}
		for i,v :=range pow{
			fmt.Printf("\n %d slice elemanları : %d",i,v )
		}

	*/
	//array
	/*
		a := [...] string {"1","2","3","4","5","6"}
		for i :=range a{
			fmt.Println("\n %d ninci eleman ",i,a[i] )
		}
	*/
	//map farklı kullanımları verilmiş tir
	baskentler := map[string]string{"türkiye": "istanbul", "fransa": "italya", "japonyua": "çin"}
	for key := range baskentler {
		fmt.Println("ülke", key, "başkent", baskentler[key])
	}
	fmt.Println("------------------------")

	for key, value := range baskentler {
		fmt.Println("ülke", key, "başkent", value)
	}

}
