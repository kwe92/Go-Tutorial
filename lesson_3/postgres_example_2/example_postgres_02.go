package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Student represents a student record.
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
		TABLE0   = "students"
	)

	student0 := Student{
		name:        "shikamaru",
		roll_number: 1001,
	}

	student1 := Student{
		name:        "Marcus Aurelius",
		roll_number: 121,
	}

	// dynamic string iterpolation.
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USERNAME, PASSWORD, DBNAME)

	// static hard coded insert statement.
	insertStatement0 := `insert into students (name, roll_number) values ('Gaara', 9999);`

	// dynamic insert statement.
	insertStatement1 := `insert into students (name, roll_number) values ($1, $2);`

	// update statement
	updateStatement0 := `update students set name = $1, roll_number = $2 where roll_number = $3;`

	// Opens database connection.
	db, err := sql.Open("postgres", psqlConn)

	CheckError(err)

	err = db.Ping()

	CheckError(err)

	fmt.Println("\nsuccessful connection!")

	_, err = db.Exec(TruncateTable(TABLE0))

	CheckError(err)

	fmt.Println("\nSuccessful truncation of table:", TABLE0)

	_, err = db.Exec(insertStatement0)

	CheckError(err)

	fmt.Println("\nsuccessful static insert!")

	_, err = db.Exec(
		insertStatement1,
		student0.name,
		student0.roll_number,
	)

	CheckError(err)

	fmt.Println("\nsuccessful dynamic insert!")

	_, err = db.Exec(updateStatement0, student1.name, student1.roll_number, student0.roll_number)

	CheckError(err)

	fmt.Printf("\nupdated record %d successfully!", student0.roll_number)

	// closes previously opened database conection.
	db.Close()

	err = db.Ping()

	CheckClosed(err)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckClosed(err error) {
	if err != nil {
		fmt.Printf("\n\nDatabase closed successfully\n\n")
		return
	}
	fmt.Printf("\nConnection is still open")

}

func TruncateTable(table_name string) string {
	return fmt.Sprintf("truncate table %s;", table_name)
}
