/*package main

import (
	"database/sql"
	"fmt"
	_"database/sql"
	_"log"
	_"github.com/go-sql-driver/mysql"
	_ "github.com/denisenkom/go-mssqldb"

)*/
package main

import (
	"database/sql"
	. "github.com/denisenkom/go-mssqldb"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mssql"
)

var (
	server   = "localhost"
	port     = 1433
	user     = "DESKTOP-80MFS0S"
	password = ""
	database = "ogrenci2"
)

func main() {
	db, err := sql.Open("mssql", "root:DESKTOP-80MFS0S@/ogrenci2")
	if err != nil {
		panic(err.Error())
	}
	db.Ping()
	defer db.Close()
}
