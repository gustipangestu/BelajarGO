package jobHandler

import (
	"encoding/json"
	"io"
	"net/http"
)

type Job struct {
	Departmen string
	Position  string
	Year      int
}

func JobGet(w http.ResponseWriter, r *http.Request) {
	myJob := Job{
		Departmen: "Malang",
		Position:  "Manager",
		Year:      10,
	}

	str, err := json.Marshal(myJob)
	if err != nil {
		io.WriteString(w, string(err.Error()))
	}
	io.WriteString(w, string(str))
}
