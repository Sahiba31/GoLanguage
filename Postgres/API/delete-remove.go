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
	get_email := r.URL.Query().Get("email")
	sqlStatement := `
	DELETE FROM students
	WHERE email = $1;`

	count, err := db.Exec(sqlStatement, get_email)
	cnt, _ := count.RowsAffected()
	if err != nil {
		fmt.Println("Error...")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if cnt == 0 {
		fmt.Println("Email is not valid...")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Println("No Error")
	defer db.Close()
}

func main() {
	http.HandleFunc("/remove", POSTHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
