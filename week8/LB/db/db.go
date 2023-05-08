package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func simpan(db *sql.DB) {
	simpan, err := db.Query("INSERT INTO lb_user (user_nama, user_email) VALUES ('Gusti Pangestu', 'gusti.pangestu@binus.ac.id')")
	if err != nil {
		panic(err.Error())
	} else {
		simpan.Close()
	}
}

func delete(db *sql.DB) {
	del, err := db.Query("DELETE FROM lb_user WHERE user_id=1")
	if err != nil {
		panic(err.Error())
	} else {
		del.Close()
	}
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lb_golang")

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("berhasil konek")
	}

	// simpan(db)
	delete(db)

}
