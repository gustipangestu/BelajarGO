package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/golang-jwt/jwt"
)

type Message struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

var sampleSecretKey = []byte("SecretYouShouldHide")
var api_key = "1234"

func CreateJwt() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	tokenStr, err := token.SignedString(sampleSecretKey)
	cekError(err)
	return tokenStr, nil
}

func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("not authorized"))
				}
				return sampleSecretKey, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("not authorized: " + err.Error()))
			}

			if token.Valid {
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not authorized"))
		}
	})
}

func getJwt(w http.ResponseWriter, r *http.Request) {
	if r.Header["Access"] != nil {
		if r.Header["Access"][0] == api_key {
			token, err := CreateJwt()
			if err != nil {
				return
			}
			fmt.Fprintf(w, token)
		}
	}
}

func cekError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func aksesHalaman(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	cekError(err)
	err = json.NewEncoder(w).Encode(msg)
	cekError(err)
}

func tes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string("hallo"))
}

func main() {

	http.Handle("/api", ValidateJWT(tes))
	http.HandleFunc("/jwt", getJwt)

	http.ListenAndServe(":3500", nil)

}
