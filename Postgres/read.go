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

type Student struct {
	ID     int
	Name   string
	Course string
	Email  string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `SELECT * FROM students WHERE student_id=$1;`
	var student Student
	var id int
	fmt.Println("Enter Id:")
	fmt.Scanln(&id)
	row := db.QueryRow(sqlStatement, id)
	err = row.Scan(&student.ID, &student.Name, &student.Course, &student.Email)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(student)
	default:
		panic(err)
	}
}
