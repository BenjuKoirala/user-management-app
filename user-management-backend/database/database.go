package database

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"

	_ "github.com/lib/pq"
)

//var DB *sql.DB

var (
	DB   *sql.DB
	PSQL = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
)

func ConnectDB() {
	connStr := "user=ashokadhikari dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	DB = db
	fmt.Println("Successfully connected to the database")
}
