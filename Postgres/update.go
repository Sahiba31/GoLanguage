package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mypass"
	dbname   = "sahibadb"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
UPDATE students
SET student_name = $2, course = $3
WHERE student_id = $1;`
	var id int
	var name, course string
	fmt.Print("Enter your id, name, and course: ")
	fmt.Scanln(&id, &name, &course)

	_, err = db.Exec(sqlStatement, id, name, course)
	if err != nil {
		panic(err)
	}
}
