package main

import "fmt"

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

func main() {
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
	fmt.Println(enes.Soyad)

}
