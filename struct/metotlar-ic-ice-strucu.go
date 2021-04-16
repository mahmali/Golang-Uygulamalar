package main

import (
	"fmt"
	_ "strconv"
)

func main() {
	//kullanici veri oluşturma alanı v1
	fmt.Println("kullanici oluşturma v1")
	kullanici1 := &kullanici{
		Id:           1,
		ad:           "muhammed",
		soyad:        "akkaya",
		kullaniciAdi: "mahmali",
		yas:          24,
		ödeme: &ödemeyapisi{
			fiyat: 20.2,
			bonus: 30.1,
		},
	}
	fmt.Println(kullanici1, kullanici1.ödeme.fiyat+kullanici1.ödeme.bonus)
	fmt.Println(kullanici1.gettümisim())
	fmt.Println(kullanici1.getödeme())
	fmt.Println("kullanıci ödeme toplamı %v", kullanici1.getödeme())

	//kullanici oluşturma v2
	kullanici2 := yenikullanici()
	kullanici2.Id = 10
	kullanici2.ad = "ahmet"
	kullanici2.soyad = "akkaya"
	kullanici2.ödeme = &ödemeyapisi{fiyat: 2.56, bonus: 654.789}
	kullanici2.yas = 55
	kullanici2.kullaniciAdi = "mahmali2"

	fmt.Println(kullanici2.getkullaniciadi())
	fmt.Println(kullanici2.getödeme())
	fmt.Println(kullanici2.gettümisim())
}

//kullanıcı yapısı
type kullanici struct {
	Id           int
	ad           string
	soyad        string
	kullaniciAdi string
	yas          int
	ödeme        *ödemeyapisi //struct içinde struct kullanma işlemi
}

//ödeme yapisi
type ödemeyapisi struct {
	fiyat float64
	bonus float64
}

//kullanicinin yapıcı metodu
func yenikullanici() *kullanici {
	u := new(kullanici)
	u.ödeme = yeniödemeyapisi()
	return u
}

//ödemenin yapıcı metodu
func yeniödemeyapisi() *ödemeyapisi {
	p := new(ödemeyapisi)
	return p
}

//metodlar
func (u kullanici) gettümisim() string {
	return u.ad + " " + u.soyad
}
func (u *kullanici) getkullaniciadi() string {
	return u.kullaniciAdi
}
func (u *kullanici) getödeme() float64 {
	öde := u.ödeme.fiyat + u.ödeme.bonus
	return öde
}
