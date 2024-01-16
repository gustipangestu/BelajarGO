package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
)

var KeyRahasia = []byte("KeySuperRahasia")
var login = "123"

func cekError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func buatToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	tokenString, err := token.SignedString(KeyRahasia)
	cekError(err)
	return tokenString, nil
}

func getJwt(w http.ResponseWriter, r *http.Request) {
	if r.Header["Access"] != nil {
		if r.Header["Access"][0] == login {
			tokenStr, err := buatToken()
			cekError(err)
			fmt.Fprintf(w, tokenStr)
		}

	}
}

func Middlewaree(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("tidak terautorizasi"))
				}
				return KeyRahasia, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("tidak terautorizasi"))
			}

			if token.Valid {
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("tidak terautorizasi"))
		}
	})
}

func akses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string("ini adalah halaman yang kamu akses"))
}

func main() {
	http.HandleFunc("/aksestanpaMW", akses)
	http.Handle("/aksesdenganMW", Middlewaree(akses))
	http.HandleFunc("/getToken", getJwt)

	http.ListenAndServe(":5050", nil)
}
