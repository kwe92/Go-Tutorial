package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Student struct {
	name        string
	roll_number int
}

func main() {
	const (
		HOST     = "localhost"
		PORT     = 5432
		USERNAME = "postgres"
		PASSWORD = "postgres"
		DBNAME   = "db_1"
	)

	// dynamic string interpolations
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USERNAME, PASSWORD, DBNAME)

	// static hard coded insert statement
	insertStatement0 := `insert into students (name, roll_number) values ('Gaara', 9999);`

	// dynamic insert statement
	insertStatement1 := `insert into students (name, roll_number) values ($1, $2);`

	db, err := sql.Open("postgres", psqlConn)

	student := Student{
		name:        "shikamaru",
		roll_number: 1001,
	}

	CheckError(err)

	err = db.Ping()

	CheckError(err)

	fmt.Println("\nsuccessful connection!")

	_, err = db.Exec(insertStatement0)

	CheckError(err)

	fmt.Println("\nsuccessful static insert!")

	_, err = db.Exec(
		insertStatement1,
		student.name,
		student.roll_number,
	)

	CheckError(err)

	// close the conection
	db.Close()

	err = db.Ping()

	CheckError(err)

	fmt.Println("successful conection!")

	CheckError(err)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
