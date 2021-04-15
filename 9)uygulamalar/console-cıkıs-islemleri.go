package main

import "fmt"

func main() {
	str1 := "muhammedali"
	str2 := "ali"
	str3 := "akkaya"
	aNumber := 32
	isTrue := true

	//varilen degerlerin uzunluklarını degişkene atma
	stringLength, _ := fmt.Println(str1, str2, str3)
	//if err == nil{
	fmt.Println("verilen stringlerin toplam uzunlugu", stringLength)
	//}
	fmt.Print("veri int %v \n", aNumber)
	fmt.Print("veri boolean \n %v", isTrue)
	fmt.Print("veri cast edilmiş %.2f \n", float32(aNumber))

	fmt.Print("data types: %T,%T,%T,%T, and %T", str1, str2, str3, aNumber, isTrue)

	//verilen degerlerin tiplerini bulmak için kullanılır
	degisken := fmt.Sprintf("data types: %T,%T,%T,%T, and %T", str1, str2, str3, aNumber, isTrue)
	fmt.Print(degisken)
}
