package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// "fmt"
// "database/sql"
//

func insert(db *sql.DB) {
	insert, err := db.Query("INSERT INTO biodata (bio_nama, bio_ttl, bio_email) VALUES ('Pangestu', '1995-05-05', 'gusti.pangestu@binus.edu')")
	if err != nil {
		panic(err.Error())
	} else {
		insert.Close()
	}
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_latihan")

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("konek")
	}

	defer db.Close()
	insert(db)

}
