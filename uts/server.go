package main

import (
	bio "biodataHandler"
	"encoding/json"
	"io"
	job "jobHandler"
	"net/http"
)

type Address struct {
	Street string
	City   string
}

type Person struct {
	Name    string
	Address Address
}

func detailDiri(w http.ResponseWriter, r *http.Request) {
	nama := r.FormValue("Nama")
	data := bio.Nama(nama)

	str, err := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		io.WriteString(w, string(err.Error()))
	} else {
		io.WriteString(w, string(str))
	}

}

func getAlldata(w http.ResponseWriter, r *http.Request) {
	var data = bio.Alldata()
	w.Header().Set("Content-Type", "application/json")
	str, err := json.Marshal(data)
	if err != nil {
		io.WriteString(w, string(err.Error()))
	} else {
		io.WriteString(w, string(str))
	}
}

func main() {
	multiplexer := http.NewServeMux()
	multiplexer.HandleFunc("/nama", detailDiri)
	multiplexer.HandleFunc("/job", job.JobGet)
	multiplexer.HandleFunc("/semuadata", getAlldata)

	http.ListenAndServe(":5050", multiplexer)

}
