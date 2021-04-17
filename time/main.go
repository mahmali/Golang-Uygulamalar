package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("şu anki unix zaman:", time.Now().Unix())
	//time.Sleep(2*time.Second)  //2 sn bekleme amaçlı yazılmıştır
	fmt.Println("şu anki unix zaman:", time.Now().Unix())

	t := time.Date(2020, time.November, 10, 20, 0, 0, 0, time.UTC) //belirli bir tarih berlirtme işlemi
	fmt.Printf("calisiyor:%s", t)

	now := time.Now() //mevcut zamannı alma işlemi
	fmt.Printf("\nmevcut zaman %s", now)

	fmt.Println("\n\n ay:", now.Month())
	fmt.Println("gün:", now.Day())
	fmt.Println("hafta:", now.Weekday())

	yarın := now.AddDate(0, 0, 1) //bir gün ekleme işlemi
	fmt.Printf("yarın: %v ,%v ,%v,%v \n", yarın.Weekday(), yarın.Month(), yarın.Day(), yarın.Year())

	uzunformat := "monday,january 2,2020" //kendi yazım formatınızı belirleyebilirsiniz
	fmt.Println("yarın", yarın.Format(uzunformat))

	tarih := time.Date(1071, 10, 11, 20, 23, 0, 0, time.UTC)
	kısaformat := "1/2/2020"
	fmt.Println("yarın", yarın.Format(kısaformat))
	//başka bir egitim videosuna geçildi
	yaz := fmt.Println
	yaz("\n", now)
	yaz("YIL:", now.Year())
	yaz("ay:", now.Month())
	yaz("gün", now.Day())
	yaz("saat:", now.Hour())
	yaz("nano saniye:", now.Nanosecond())
	yaz("lokasyon:", now.Location())
	yaz("dakika:", now.Minute())
	yaz("saniye:", now.Second())

	yaz(now.Weekday())
	//tarih karşılaştırma yada test yada kontrol yapma işlemi
	yaz("\n", tarih.Before(now)) //tarih şu anki tarihten öncemi
	yaz(tarih.After(now))        //tarih şu anki tarihten sonramı
	yaz(tarih.Equal(now))        //varilen tarihler eşitmi
	yaz("\n")
	fark := now.Sub(tarih) //tarihler arasındaki farkı verir
	yaz(fark)

	yaz(fark.Hours())
	yaz(fark.Minutes())
	yaz(fark.Seconds())

}
