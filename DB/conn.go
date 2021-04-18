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

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	res, _ := db.Exec("insert into musteri(id,ad,soyad,sehir) values (5,'ali','akkaya','maras')")
	//db.Exec("insert into musteri(id,ad,soyad,sehir) values (4,'bismillah','akkaya','maras')")
	rowCount, _ := res.LastInsertId()
	log.Printf("etkilenen kayıt sayısı", rowCount)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
