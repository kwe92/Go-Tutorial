package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// Student represents a student record.
type Student struct {
	id           int
	name         string
	roll_number  int
	created_date string
}

func main() {

	// enum of constants.
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

	student2 := Student{
		name:        "Carl Jung",
		roll_number: 1961,
	}

	// dynamic string iterpolation.
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USERNAME, PASSWORD, DBNAME)

	// static hard coded insert statement.
	insertStatement0 := `insert into students (name, roll_number, created_date) values ('Gaara', 9999, '2023-09-04');`

	// dynamic insert statement.
	insertStatement1 := `insert into students (name, roll_number, created_date) values ($1, $2, $3);`

	insertStatement2 := `insert into students (name, roll_number, created_date) values ($1, $2, $3);`

	// update statement.
	updateStatement0 := `update students set name = $1, roll_number = $2, created_date = $3 where roll_number = $4;`

	// delete statement.
	deleteStatement0 := `delete from students where roll_number = $1;`

	// select query statement.

	selectAll := `select * from students;`

	// Slice of students

	students := make([]Student, 0)

	// Opens database connection.
	db, err := sql.Open("postgres", psqlConn)

	CheckError(err)

	// ping database to ensure connection is successful.
	err = db.Ping()

	CheckError(err)

	fmt.Println("\nsuccessful connection!")

	// clear table of data.
	_, err = db.Exec(TruncateTable(TABLE0))

	CheckError(err)

	fmt.Println("\nSuccessful truncation of table:", TABLE0)

	// execute static insert statement.
	_, err = db.Exec(insertStatement0)

	CheckError(err)

	fmt.Println("\nsuccessful static insert!")

	// execute dynamic insert statement with additional interpolated arguments.
	_, err = db.Exec(
		insertStatement1,
		student0.name,
		student0.roll_number,
		GetDate(),
	)

	CheckError(err)

	_, err = db.Exec(
		insertStatement2,
		student2.name,
		student2.roll_number,
		GetDate(),
	)

	CheckError(err)

	fmt.Println("\nsuccessful dynamic insert!")

	// execure update statement with interpolated arguments.
	_, err = db.Exec(
		updateStatement0,
		student1.name,
		student1.roll_number,
		GetDate(),
		student0.roll_number,
	)

	CheckError(err)

	fmt.Printf("\nupdated record %d successfully!", student0.roll_number)

	// execute delete statement with interpolated arguments.
	_, err = db.Exec(deleteStatement0, 9999)

	CheckError(err)

	fmt.Printf("\n\ndeleted record: %d successfully!", 9999)

	// ?---------------SELECT DATA AND ENUMERATE THROUGH ROWS---------------? //

	rows, err := db.Query(selectAll)
	CheckError(err)

	// close iteration of rows when the program ends.
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var roll_number int
		var created_date string

		err = rows.Scan(
			&id,
			&roll_number,
			&name,
			&created_date,
		)

		CheckError(err)

		fmt.Printf("\n\nroll_number:%d\nname: %s", roll_number, name)
	}

	// ?------------------ROW ENUMERATION WITH STRUCT------------------? //

	rows, err = db.Query(selectAll)

	for rows.Next() {
		var student Student

		err = rows.Scan(
			&student.id,
			&student.roll_number,
			&student.name,
			&student.created_date,
		)

		CheckError(err)

		fmt.Printf("\n\nstudent: %v", student)

		students = append(students, student)

	}

	fmt.Printf("\n\nstudents slice: %v", students)

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
		fmt.Printf("\n\nDatabase closed successfully!\n\n")
		return
	}
	fmt.Printf("\nConnection is still open!")

}

func TruncateTable(table_name string) string {
	return fmt.Sprintf("truncate table %s;", table_name)
}

func GetDate() string {
	return time.Now().Format("2006-01-02")
}
