package main

import "fmt"

func main() {
	küp, kare := islem(9)
	fmt.Println(küp, kare)
}
func islem(deger int) (x, y int) {
	x = deger * deger * deger
	y = deger * deger
	return
}
