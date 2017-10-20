package model

import (
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "login success")
}

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "logout success")
}
