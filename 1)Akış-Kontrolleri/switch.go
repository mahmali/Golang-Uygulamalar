package main

import (
	"fmt"
)

func main() {
	/*
		foo:=0

		switch {
		case foo==1:
			println("sayı degeri 1")
		case foo==2:
			println("sayi 2 ye eşittir")
		case foo>2:
			println("sayi 2 den dahada büyüktür")
		default:
			println("verien degerler arasında bir deger verilemedi")
		}
	*/
	var score float64

	for { //for burada sonsuz döngü olarak kullanılmatadır
		fmt.Println("console dan bir deger giriniz")
		fmt.Scanf("%v", &score) //consoleden deger girdirmek için yapılmıştır

		switch {
		case score <= 59:
			println("FF")
		case score < 70:
			println("deger 60 70 arasındadır", score)
		case score < 80:
			println("deger 70 80 arasındadır")
		case score > 80:
			println("deger 80 den büyük bir degerdir")

		}
	}

}
