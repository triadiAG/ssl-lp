package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Response(w, http.StatusOK, "hello")
	}).Methods(http.MethodGet)
	r.HandleFunc("/.well-known/pki-validation", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("A23AB2849F0881461F4840FF2256654F7F3FD0934549DB5ACF70A4665370E32F\ncomodoca.com\nff89043f12fade2"))
	}).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":7778", r))
}

func Response(w http.ResponseWriter, status int, body interface{}) {
	payload, _ := json.Marshal(body)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	w.Write(payload)
}
