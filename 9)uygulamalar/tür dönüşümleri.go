package main

import (
	"fmt"
	"strconv"
)

func main() {
	//convert işlemleri
	var myString = "1"
	var x = 10
	var f float32 = 2.0

	fmt.Println(myString, x, f)

	//string den int e dösnüşüm yapma işlemi
	number, _ := strconv.Atoi(myString)
	fmt.Println(number)

	result := number + 2
	fmt.Println(result)

	fmt.Println("sonuc = " + strconv.Itoa(number))

	var i int = 55
	var fl float64 = float64(i)
	var u uint = uint(fl)

	fmt.Println(i, fl, u)

}
