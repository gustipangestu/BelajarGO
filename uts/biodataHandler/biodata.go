package biodataHandler

import "strings"

type Bio struct {
	Nama   string
	Nim    string
	Alamat string
}

var data1 = Bio{
	Nama:   "Gusti",
	Nim:    "1234",
	Alamat: "Malang",
}

var data2 = Bio{
	Nama:   "Pangestu",
	Nim:    "321",
	Alamat: "Kabupaten Malang",
}

func Nama(nama string) Bio {
	// io.WriteString(w, "Hello nama saya adalah Mr. X")

	// w.Header().Set("Content-Type", "application/json")
	var data Bio
	if strings.EqualFold(nama, "gusti") {
		data = data1
	} else if strings.EqualFold(nama, "pangestu") {
		data = data2
	}
	return data

}

func Alldata() []Bio {
	var data = []Bio{data1, data2}
	return data
}
