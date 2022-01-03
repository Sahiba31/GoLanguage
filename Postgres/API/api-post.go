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
	Name   string `json:"student_name"`
	Email  string `json:"email"`
	Course string `json:"course"`
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

func POSTHandlerCreate(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()
	var s Student
	err := json.NewDecoder(r.Body).Decode(&s)
	fmt.Println(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error....")
		return
	}
	sqlStatement := `INSERT INTO student (student_name, email, course)
	VALUES ($1, $2, $3)
	RETURNING student_id`
	student_id := 0

	err = db.QueryRow(sqlStatement, s.Name, s.Email, s.Course).Scan(&student_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Id of the newly added student is: ", student_id)
	defer db.Close()
}

func GETHandler(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()
	get_id := r.URL.Query().Get("id")
	sqlStatement := `SELECT student_name, email FROM students WHERE student_id=$1;`
	var student Student

	row := db.QueryRow(sqlStatement, get_id)
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

func POSTHandlerUpdate(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()
	var s Student
	err := json.NewDecoder(r.Body).Decode(&s)
	fmt.Println(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sqlStatement := `
	UPDATE students
	SET student_name = $2
	WHERE email = $1;`

	count, err := db.Exec(sqlStatement, s.Email, s.Name)
	cnt, _ := count.RowsAffected()
	// fmt.Println(count)
	if err != nil {
		// fmt.Println("Error...")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Hiiii"))
		return
	}
	if cnt == 0 {
		fmt.Println("Invalid Email")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Hiiii"))
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Println("No Error")
	defer db.Close()
}

func POSTHandlerDelete(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()
	var s Student
	err := json.NewDecoder(r.Body).Decode(&s)
	fmt.Println(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error....")
		return
	}
	sqlStatement := `
	DELETE FROM students
	WHERE email = $1;`

	count, err := db.Exec(sqlStatement, s.Email)
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
	response := http.NewServeMux()
	response.HandleFunc("/create", POSTHandlerCreate)
	response.HandleFunc("/fetch", GETHandler)
	response.HandleFunc("/update", POSTHandlerUpdate)
	response.HandleFunc("/remove", POSTHandlerDelete)
	servers := http.Server{Addr: ":8080", Handler: response}
	servers.ListenAndServe()
}
