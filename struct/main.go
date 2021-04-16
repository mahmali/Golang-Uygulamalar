package main

import (
	"fmt"
	"strconv"
)

type Insan struct {
	Isim  string
	Soyad string
	Yas   int
}

//varsayılan ve boş yapıcı metod
func yeniInsan() *Insan {
	i := new(Insan)
	return i
}
func yeniInsanParametre(isim, soyisim string, yas int) *Insan {
	h := new(Insan)
	h.Isim = isim
	h.Soyad = soyisim
	h.Yas = yas
	return h
}

func main() { /*
		ali := Insan{Isim: "muhammed ", Soyad: "akkaya ", Yas: 24}
		fmt.Print(ali.Isim)
		fmt.Print(ali.Soyad)
		fmt.Println(ali.Yas)

		veli := new(Insan)
		veli.Isim = "veli"
		veli.Soyad = "nacar"
		veli.Yas = 35
		fmt.Println(veli.Isim, veli.Soyad, veli.Yas)

		enes := yeniInsan()
		enes.Soyad = "pazar"
		fmt.Println(enes.Soyad)*/

	varlık := yeniInsanParametre("ahmet", "zahit", 15) //burada 2 tanesi string 1 tane int oldugundan hata verdi
	//fmt.Println(varlık.Isim,varlık.Soyad,varlık.Yas)
	var data = varlık.Isim + " " + varlık.Soyad + " " + strconv.Itoa(varlık.Yas) //hatanın önüne geçebilmek için int string veri tipine dönüştürüldü
	fmt.Println(data)

}
