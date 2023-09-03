package main

import (
	"database/sql"
	"fmt"

	// "net/http"
	// "log"
	_ "github.com/lib/pq"
)

// database structure that handles one or more connections
var db *sql.DB

// This function will make a connection to the database only once.
func init() {

	var err error

	connStr := "postgres://postgres:password@localhost/db_1?sslmode=disable"

	// opens a database specified by its database driver.

	db, err = sql.Open("postgres", connStr)

	CheckError(err)

	// check if database is alive, if err is not null then the connection is not alive.

	err = db.Ping()

	CheckError(err)

	// closes a database.
	defer db.Close()

	// check if connection is stil alive.
	err = db.Ping()

	CheckError(err)

	fmt.Println("\nThe database is connected.\n ")

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// fmt.Println("from main function")

}

// init function

//   - initialization function, runs at the initial start of your application
//   - init can be declared multiple times, executed in the order defined

// sql package: Standard Library

//   - Provides a generic interface through a struct for SQL and SQL-lite databases
//   - requires a third party database driver that works with GO (GO provides a list of drivers within their documentation)
//   -
