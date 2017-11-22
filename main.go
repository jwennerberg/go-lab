package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("go api testing\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api", ApiHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}
