package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"fmt"
)

func main() {
	r := mux.NewRouter()
	r.Path("/{category}").
		HandlerFunc(HomeHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func Middleware(h http.Handler) http.Handler {

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Println(vars)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}