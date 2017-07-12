package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

// vim: ts=4 sw=4 expandtab :
