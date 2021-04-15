package main

import "fmt"

func main() {
	dizi := [4]int{1, 2, 3, 4}
	dizi2 := [...]string{"muhammed", "ali", "akkaya"}

	for _, eleman := range dizi {
		fmt.Println(eleman)
	}
	for _, eleman := range dizi2 {
		fmt.Println(eleman)
	}
}
