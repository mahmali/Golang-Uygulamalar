package main

import "fmt"

func main() {
	dilim := [...]int{1, 2, 3}
	dilim1 := dilim[:]
	fmt.Println(dilim, dilim1)
	dilim1[0] = 10
	fmt.Println(dilim, dilim1) //dilmler pointer mantıgı ile çalıştıgından iki dizidede deger degişmiştir
}
