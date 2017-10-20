package model

import (
	"fmt"
	"net/http"
)

func Userinfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "login success")
}
