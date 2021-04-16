package main

import (
	"fmt"
)

func main() {
	//kullanici veri oluşturma alanı v1
	fmt.Println("kullanici oluşturma v1")
	kullanici := &kullanici{
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
	fmt.Println(kullanici)
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
