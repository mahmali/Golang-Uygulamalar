package main

import (
	"fmt"
	"time"
)

func main() {
	// date() metodu ile kendi tarih verimi oluşturuyoruz.
	t := time.Date(2016, time.November, 10, 20, 0, 0, 0, time.UTC)

	//t adıylı tarih tipinde oluşturulan veriyi string tipinde ekrana yazıyoruz.
	fmt.Println("tarih: ", t)

	//timw.noe() ile anında zaman bilgisi alma işlemi yapılır
	now := time.Now()
	fmt.Println(now)

	// başta oluşturulan zaman bilgisi ile ay gün yıl ı ayrı bir şekilde alma işlemi yapma
	fmt.Println("gün", t.Day())
	fmt.Println("ay", t.Month())
	fmt.Println("hafta", t.Weekday())

	//tarihe 1 gün ekle
	tomorrow := t.AddDate(0, 0, 1)
	fmt.Println("yarın",
		tomorrow.Weekday(),
		tomorrow.Month(),
		tomorrow.Day(),
		tomorrow.Year())

	longFormat := "monday, junuary 2,2006 "
	fmt.Println("tomorrow", tomorrow.Format(longFormat))

	shortFormat := "1/2/06"
	fmt.Println("yarın", tomorrow.Format(shortFormat))
}
