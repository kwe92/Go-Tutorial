package main

import (
	"database/sql"

	"fmt"

	_ "github.com/lib/pq"
)

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "postgres"
	DBNAME   = "db_1"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DBNAME)

	// open database.
	db, err := sql.Open("postgres", psqlconn)

	CheckError(err)

	err = db.Ping()

	CheckError(err)

	// close database
	defer db.Close()

	err = db.Ping()

	CheckError(err)

	fmt.Println("Conneted to the database successfully!")

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
