package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lb_golang")

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("berhasil konek")
	}

	db.Close()
}
