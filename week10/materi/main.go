package main

import (
	"Mhs"
	mhs "Mhs"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func tampilMhs(w http.ResponseWriter, r *http.Request) {
	var mhs = mhs.SelectAll()
	str, err := json.Marshal(mhs)
	if err != nil {
		panic(err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, string(str))
	}
}

func simpanMhs(w http.ResponseWriter, r *http.Request) {
	nama := r.FormValue("nama")
	nim := r.FormValue("nim")
	nimm, _ := strconv.ParseInt(nim, 10, 64)
	email := r.FormValue("email")
	hp := r.FormValue("no_hp")

	Mhs.Simpan(int(nimm), nama, email, hp)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/tampilMhs", tampilMhs)
	mux.HandleFunc("/simpanMhs", simpanMhs)

	http.ListenAndServe(":5050", mux)
}
