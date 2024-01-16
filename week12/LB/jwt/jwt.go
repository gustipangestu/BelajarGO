package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	_ "github.com/golang-jwt/jwt"
)

var kodeRahasia = []byte("BINUSMalang")
var api_key = "123456789"

func CreateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	tokenStr, err := token.SignedString(kodeRahasia)
	if err != nil {
		panic(err.Error())
	}
	return tokenStr, nil
}

func GetJWT(w http.ResponseWriter, r *http.Request) {
	if r.Header["Access"] != nil {
		if r.Header["Access"][0] == api_key {
			token, err := CreateJWT()
			if err != nil {
				return
			}
			fmt.Fprintf(w, token)
		}
	}
}

// middleware
func ValidasiJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			fmt.Print("ada tokennya")
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, oke := t.Method.(*jwt.SigningMethodHMAC)
				if !oke {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("tidak terauthorized"))
				}
				return kodeRahasia, nil
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("tidak terauthorized"))
			}

			if token.Valid {
				next(w, r)
			}
		}
	})
}

func Halaman1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string("Hello"))
}

func Tes1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string("Hello Tes1"))
}

func main() {
	http.Handle("/halaman1", ValidasiJWT(Halaman1))
	http.HandleFunc("/getJwt", GetJWT)
	http.Handle("/tes", ValidasiJWT(Tes1))

	http.ListenAndServe(":5050", nil)

}
