package model

import (
	"fmt"
	"net/http"
)

type User struct{
	area_code string
	mobile string
	email string

}


func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "login success")
}

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "logout success")
}
