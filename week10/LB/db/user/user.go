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

func cekError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func koneksi() sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/la_golang")
	cekError(err)
	defer db.Close()
	return *db
}

func SelectAll() []User {
	db := koneksi()
	result, err := db.Query("SELECT * FROM la_user")
	cekError(err)
	var user = User{}
	var userArray []User
	for result.Next() {
		err = result.Scan(&user.Nim, &user.Nama, &user.Email, &user.Hp)
		cekError(err)
		userArray = append(userArray, User{Nim: user.Nim, Nama: user.Nama, Email: user.Email, Hp: user.Hp})
	}
	for i := 0; i < len(userArray); i++ {
		fmt.Println(userArray[i])
	}
	return userArray
}
