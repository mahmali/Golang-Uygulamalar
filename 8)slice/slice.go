package main

import "fmt"

func main() {
	dilim := []int{1, 2, 3, 4, 5, 6}

	var s []int = dilim[1:3]

	fmt.Println(dilim)
	fmt.Println(s)

	fmt.Println("kullanılar dilimin alanı:", cap(dilim), "boyutu", len(dilim))

}
