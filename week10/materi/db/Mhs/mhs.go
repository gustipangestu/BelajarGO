package Mhs

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Mhs struct {
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

func konek() sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/la_golang")
	cekError(err)
	defer db.Close()
	return *db
}

func SelectAll() []Mhs {
	db := konek()

	statement, err := db.Query("SELECT * FROM la_user")
	cekError(err)
	var mhs = Mhs{}
	var mhsArray []Mhs
	for statement.Next() {
		err = statement.Scan(&mhs.Nim, &mhs.Nama, &mhs.Email, &mhs.Hp)
		cekError(err)
		mhsArray = append(mhsArray, Mhs{Nim: mhs.Nim, Nama: mhs.Nama, Email: mhs.Email, Hp: mhs.Hp})
	}
	for i := 0; i < len(mhsArray); i++ {
		fmt.Println(mhsArray[i])
	}
	return mhsArray
}

func Simpan(nim int, nama, email, hp string) {
	db := konek()
	statement, err := db.Prepare("INSERT INTO la_user(nim, nama, email, no_hp) VALUES (?,?,?,?)")
	cekError(err)
	respon, err := statement.Exec(nim, nama, email, hp)
	cekError(err)
	id, err := respon.LastInsertId()
	cekError(err)
	fmt.Print(id)
}
