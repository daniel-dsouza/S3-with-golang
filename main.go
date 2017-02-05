package main

import (
	"log"
	"net/http"

	"github.com/daniel-dsouza/test/app/bundles/uploadphotos"
	"github.com/gorilla/mux"
)

func main() {
	//Controllers Declaration

	photocontroller := &uploadphotos.UploadPhotoController{}

	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1/").Subrouter()

	s.HandleFunc("/photo", photocontroller.Upload).Methods("POST")
	s.HandleFunc("/down", photocontroller.Download).Methods("GET")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
