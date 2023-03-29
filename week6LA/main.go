package main

import (
	"Bangun2D"
	"Bangun3D"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func luasPersegi(w http.ResponseWriter, r *http.Request) {
	sisi := 10
	luas := Bangun2D.LuasPersegi(sisi)

	//ubah ke json
	str, err := json.Marshal(luas)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Print(err)
	} else {
		io.WriteString(w, string(str))
	}

	// io.WriteString(w, string(strconv.FormatFloat(luas, 'f', -1, 64)))
}

func volumeTabung(w http.ResponseWriter, r *http.Request) {
	jari, _ := strconv.Atoi(r.FormValue("Jari"))
	t, _ := strconv.Atoi(r.FormValue("Tinggi"))

	vol := Bangun3D.HitungVolumeTabung(jari, t)

	str, err := json.Marshal(vol)
	if err != nil {
		fmt.Print(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, string(str))
	}
}

func main() {
	var mux = http.NewServeMux()

	mux.HandleFunc("/luasPersegi", luasPersegi)
	mux.HandleFunc("/volumeTabung", volumeTabung)

	http.ListenAndServe(":5050", mux)
}
