package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func ConnectDB() {

	err := godotenv.Load("config/.env")

	if err != nil {
		log.Fatal("Fail To Get Env")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

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
