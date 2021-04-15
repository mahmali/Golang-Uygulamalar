package main

import "fmt"

func main() {
	renkler := []string{}
	renkler = append(renkler, "muhammed", "ali", "akkaya")
	fmt.Println(renkler)
	renkler = append(renkler, "24 yaşında")
	fmt.Println(renkler)
	renkler = append(renkler[1:len(renkler)])
	fmt.Println(renkler)
	renkler = append(renkler[:len(renkler)-1])
	fmt.Println(renkler)
}
