package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("postgres", os.Getenv("DB"))
	if err != nil {
		panic(err.Error())
	}
	DB = db
	fmt.Println("Successfully connected to the database.")
}
