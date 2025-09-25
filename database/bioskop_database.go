package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "bioskop_db"
)

var Db *sql.DB

func ConnectDB() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error
	Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error")
	}

	err = Db.Ping()
	if err != nil {
		log.Fatal(" Database tidak merespon:", err)
	}

	fmt.Println(" Success Connect DB")
}
