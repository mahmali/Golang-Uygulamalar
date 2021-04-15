package main

import (
	"../utils"
	"fmt"
)

//burdada farklı 2 şekilde kullanılan fonksiyonların kullanımı verilmiş durumdadır.
//fonsiyonlar başka bir pakette alınmış halde kullanılır.
func main() {
	ad, soyad := utils.FullName("muhammed ali", "akkaya")
	fmt.Println(ad, soyad)

	adı, soyadı := utils.FullNameLen("muhammed ali", "akkaya")
	fmt.Println(adı, soyadı)
}
