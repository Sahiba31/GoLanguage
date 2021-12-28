package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Student struct {
	Name  string
	Email string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mypass"
	dbname   = "sahibadb"
)

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

func GETHandler(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	sqlStatement := `SELECT student_name, email FROM students WHERE student_id=$1;`
	var student Student
	var id int
	fmt.Print("Enter Id: ")
	fmt.Scanf("%d", &id)
	fmt.Println()
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&student.Name, &student.Email)
	var studentBytes []byte
	if err != nil {
		studentBytes, _ = json.MarshalIndent("There is no such record for this ID...", "", "\t")
		w.Header().Set("Content-Type", "application/json")
		w.Write(studentBytes)
		log.Fatal(err)
	} else {
		studentBytes, _ = json.MarshalIndent(student, "", "\t")
		w.Header().Set("Content-Type", "application/json")
		w.Write(studentBytes)
	}

	defer db.Close()
}

func main() {
	http.HandleFunc("/", GETHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
