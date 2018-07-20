package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest/peoples"
)

func main() {
	router := mux.NewRouter()

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/intranet")
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	//Comment
	result, err := db.Query("SELECT * FROM employee LIMIT 5")
	fmt.Println(result.Columns())

	router.HandleFunc("/people", peoples.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", peoples.GetPerson).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
