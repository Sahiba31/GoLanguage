package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

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

func POSTHandler(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()
	get_name := r.URL.Query().Get("name")
	get_email := r.URL.Query().Get("email")
	get_course := r.URL.Query().Get("course")
	if get_name == "" || get_email == "" {
		fmt.Println("You should provide both name and email...")
		return
	}
	sqlStatement := `INSERT INTO student (student_name, email, course)
	VALUES ($1, $2, $3)
	RETURNING student_id`
	student_id := 0
	// var name, email, course string
	// fmt.Print("Enter your name, email, course: ")
	// fmt.Scanln(&name, &email, &course)
	err := db.QueryRow(sqlStatement, get_name, get_email, get_course).Scan(&student_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Id of the newly added student is: ", student_id)
	defer db.Close()
}

func main() {
	http.HandleFunc("/create", POSTHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
