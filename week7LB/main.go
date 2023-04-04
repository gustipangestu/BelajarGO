package main

import (
	"Pria"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func hitungBBPria(w http.ResponseWriter, r *http.Request) {
	nama := r.FormValue("Nama")
	tb := r.FormValue("Tb")
	tbfloat, _ := strconv.ParseFloat(tb, 64)
	bbIdeal := Pria.HitungBBIdealPria(nama, tbfloat)

	str, _ := json.Marshal(bbIdeal)

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(str))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hitungbbpria", hitungBBPria)

	http.ListenAndServe(":5050", mux)
}
