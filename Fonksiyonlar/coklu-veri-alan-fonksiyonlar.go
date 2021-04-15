package main

import "fmt"

func main() {
	uzunluk, toplam := ekle(1, 2, 3, 4, 5, 6)
	fmt.Printf("fonsiyona %v tane dosya gitmiş var toplamları %v dir", uzunluk, toplam)
}

func ekle(sayılar ...int) (int, int) { //fonsiyon gelen verilere listeye atarak kullanmamız saglar. belli bir sınırı yoktur.
	sayac := 0
	for _, sayı := range sayılar {
		sayac += sayı
	}
	return len(sayılar), sayac
}
