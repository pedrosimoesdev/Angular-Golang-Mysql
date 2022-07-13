package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Cars struct {
	id    int
	name  string
	model string
	year  int
}

func dbConn() (db *sql.DB) {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/crud_cars")
	db.SetConnMaxLifetime(time.Minute * 4) // <-- this
	if err != nil {
		log.Println("error")

		panic(err.Error())
	}
	return db
}

func getRecords(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM cars")
	if err != nil {
		log.Println("error 2 ")
		panic(err.Error())
	}
	car := Cars{}
	response := []Cars{}
	for selDB.Next() {
		var id int
		var name string
		var model string
		var year int
		err = selDB.Scan(&id, &name, &model, &year)
		if err != nil {
			log.Println("error 3")
			panic(err.Error())
		}
		car.id = id
		car.name = name
		car.model = model
		car.year = year
		//used to fill array
		response = append(response, car)

	}
	log.Println(response)
	fmt.Fprint(w, response)
	defer db.Close()
}

func insertRecods(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	defer db.Close()

	sql := "INSERT INTO cars(name, model,year) VALUES (?,?,?)"
	res, err := db.Exec(sql, "Opel", "test", 1997)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
		log.Fatal(lastId)
	}

	fmt.Fprintln(w, "insert Records")

}

func updateRecods(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	defer db.Close()

	sql := `UPDATE cars 
	       SET name = ? , model = ?, year = ?
	       WHERE id = ?;`
	_, err := db.Exec(sql, "opelUpdate", "modelUpdated", 2022, 6)

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintln(w, "Updated Record")
}

func deleteRecods(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	defer db.Close()

	sql := `DELETE FROM cars WHERE id = ?;`
	_, err := db.Exec(sql, 6)

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintln(w, "Deleted Record")
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", getRecords)
	http.HandleFunc("/insert", insertRecods)
	http.HandleFunc("/update", updateRecods)
	http.HandleFunc("/delete", deleteRecods)
	http.ListenAndServe(":8080", nil)
}
