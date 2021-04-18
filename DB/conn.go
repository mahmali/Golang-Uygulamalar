package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1"
	dbname   = "DBdeneme"
)

func main() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// check db
	err = db.Ping()
	CheckError(err)

	// close database
	defer db.Close()
	fmt.Println("Connected!")

	var (
		Id    int
		Ad    string
		Soyad string
		Sehir string
	)
	//kayıt ekleme satırları. her çalışmaya aynı veri eklendigi için yorum satırı yapmaya karar verdim
	/*
		res, _ := db.Exec("insert into musteri(id,ad,soyad,sehir) values (7,'ali','akkaya','maras')")
		rowCount, _ := res.LastInsertId()
		log.Printf("etkilenen kayıt sayısı", rowCount)
	*/

	/*

		//eklenen kayıtları getirme işlemi
		rows, err := db.Query("select *from musteri")
		if err != nil {
			log.Fatal(err)
		}
		//rows.columns tablodaki kolan isimlerini getirme işlemi yapmaya yarar
		for rows.Next() {
			err = rows.Scan(&Id, &Ad, &Soyad, &sehir)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("id: %v adı: %v soyadı: %v sehir: %v", Id, Ad, Soyad, Sehir)

		}
	*/

	rows, err1 := db.Query("select *from musteri where id = 1")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&Id, &Ad, &Soyad, &Sehir)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("id si %v olna veriler %v %v %v", Id, Ad, Soyad, Sehir)

	}
	err1 = rows.Err()
	if err1 != nil {
		log.Fatal(err1)
	}

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
