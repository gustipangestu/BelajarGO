package main

import (
	T "Tugas"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type DatadiriMahasiswa struct {
	Nim   string
	Nama  string
	Email string
}

func welcome(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello")
}

func hitungNilaiTugas(w http.ResponseWriter, r *http.Request) {
	nilaiTugas := T.HitungNilaiTugas(90.0, 0.4)
	io.WriteString(w, string(strconv.FormatFloat(float64(nilaiTugas), 'E', -1, 32)))
}

func inputPost(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("nama")
	umur := r.FormValue("umur")
	respon := "nama adalah " + name + ", umurnya " + umur
	str, err := json.Marshal(respon)
	if err != nil {
		io.WriteString(w, err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(str))
}

func tampilDatadiri(w http.ResponseWriter, r *http.Request) {
	dataDiri := DatadiriMahasiswa{
		Nim:   "123",
		Nama:  "Alex",
		Email: "alex123@binus.ac.id",
	}
	str, err := json.Marshal(dataDiri)
	if err != nil {
		io.WriteString(w, err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(str))

}

func no1() {
	ar1 := [...]string{"a", "c", "d"}
	ar2 := [...]string{"a", "d", "c"}

	// print array 1
	fmt.Print("array 1:\n")
	for i := 0; i < len(ar1); i++ {
		fmt.Printf("%s ", ar1[i])
	}
	fmt.Println("")

	// print array2
	fmt.Print("array 2:\n")
	for i := 0; i < len(ar1); i++ {
		fmt.Printf("%s ", ar2[i])
	}

	fmt.Println("")

	if ar1 == ar2 {
		fmt.Println("array sama")

	} else {
		for i := 0; i < len(ar1); i++ {
			hasilBanding := (ar1[i] == ar2[i])
			if hasilBanding == false {
				fmt.Printf("index ke %d berbeda \n", i)

			}
		}
	}

}

func main() {

	multiplexer := http.NewServeMux()
	multiplexer.HandleFunc("/tugas", hitungNilaiTugas)
	multiplexer.HandleFunc("/welcome", welcome)
	multiplexer.HandleFunc("/datadiri", tampilDatadiri)
	multiplexer.HandleFunc("/infut", inputPost)
	// no1()

	http.ListenAndServe(":5050", multiplexer)
}
