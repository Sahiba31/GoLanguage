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

	var students []Student
	rows, err := db.Query("SELECT * FROM students LIMIT $1", 3)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name, course, email string
		err = rows.Scan(&id, &name, &course, &email)
		if err != nil {
			panic(err)
		}
		temp := Student{id, name, course, email}
		students = append(students, temp)
	}
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(students)
	default:
		panic(err)
	}

}
