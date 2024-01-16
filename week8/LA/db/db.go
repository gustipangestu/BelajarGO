package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func insert(db *sql.DB) {
	// insert data
	insert, err := db.Query("INSERT INTO la_user (nim, nama, email, no_hp) VALUES ('123', 'Gusti', 'gusti@binus.ac.id', '08989898989')")
	if err != nil {
		panic(err.Error())
	} else {
		insert.Close()
	}
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/la_golang")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Print("berhasil konek")
		defer db.Close()
	}

	insert(db)
}
