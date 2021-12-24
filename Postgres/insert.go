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
	// fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
	INSERT INTO students (student_id, student_name, email, course)
	VALUES ($1, $2, $3, $4)`
	var id int
	var name, email, course string
	fmt.Print("Enter your id, name, email, course: ")
	fmt.Scanln(&id, &name, &email, &course)

	_, err = db.Exec(sqlStatement, id, name, email, course)
	if err != nil {
		panic(err)
	}
}
