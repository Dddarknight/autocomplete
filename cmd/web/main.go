package main

import (
	"log"
    "net/http"
)

func main() {
	mux := http.NewServeMux()
    mux.HandleFunc("/autocomplete", autocomplete)

    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}