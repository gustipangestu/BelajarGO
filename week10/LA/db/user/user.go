package user

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Nim   int    `json:"nim"`
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Hp    string `json:"no_hp"`
}

func cekEror(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func koneksi() sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/la_golang")
	cekEror(err)
	defer db.Close()
	return *db
}

func SelectAll() []User {
	db := koneksi()

	result, err := db.Query("SELECT * FROM la_user")
	cekEror(err)
	var userArray []User
	var usr = User{}

	for result.Next() {
		err := result.Scan(&usr.Nim, &usr.Email, &usr.Nama, &usr.Hp)
		cekEror(err)
		userArray = append(userArray, User{Nim: usr.Nim, Nama: usr.Nama, Email: usr.Email, Hp: usr.Hp})
	}
	fmt.Print(len(userArray))
	for i := 0; i < len(userArray); i++ {
		fmt.Println(userArray[i])
	}
	return userArray
}

func main() {
	SelectAll()
}
