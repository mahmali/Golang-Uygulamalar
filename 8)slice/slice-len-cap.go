package main

import (
	"fmt"
)

func main() {
	numara := make([]int, 5, 10)
	numara[0] = 1
	numara[1] = 2
	numara[2] = 3
	numara[3] = 4
	numara[4] = 5

	fmt.Println(numara)
	fmt.Printf("dilimin uzunluk %v boyutu %v \n", len(numara), cap(numara))
	numara = append(numara, 123)
	fmt.Println(numara) //dilime eleman eklme işlemi

}
