package main

import (
	"log"
	"net/http"
)

func main() {
	registerStaticResource()

	err := http.ListenAndServe("127.0.0.1:8800", staticMux)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func registerStaticResource() {
	http.Handle("/js/bootstrap/css/", http.FileServer(http.Dir(STATIC_FOLDER)))
	http.Handle("/js/bootstrap/js/", http.FileServer(http.Dir(STATIC_FOLDER)))
	http.Handle("/js/", http.FileServer(http.Dir(STATIC_FOLDER)))
	http.Handle("/img/", http.FileServer(http.Dir(STATIC_FOLDER)))
}
