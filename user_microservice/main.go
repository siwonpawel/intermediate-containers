package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", hello).Methods(http.MethodGet)

	http.ListenAndServe("localhost:8011", router)
}

func hello(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello user microservice!"))
}
