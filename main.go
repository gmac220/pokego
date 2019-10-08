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
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	db.Exec("INSERT INTO pokemon VALUES(4, 'Eevee')")
	getAll(db)
}

func ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected")
}

func getAll(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM POKEMON")
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
}

func searchByName(db *sql.DB, searchvalue string) {
	row := db.QueryRow("SELECT * FROM pokemon WHERE name = $1", searchvalue)
	var id int
	var name string
	row.Scan(&id, &name)
	fmt.Println(id, name)
}
