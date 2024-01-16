package main

import (
	"encoding/json"
	"io"
	"net/http"
	usr "user"
)

func cekError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func selectUser(w http.ResponseWriter, r *http.Request) {
	var data = usr.SelectAll()
	js, err := json.Marshal(data)
	cekError(err)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(js))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/selectUser", selectUser)

	http.ListenAndServe(":5050", mux)
}
