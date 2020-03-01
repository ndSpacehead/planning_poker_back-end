package main

import (
	"log"
	"net/http"

	"./engine"
)

func main() {
	server := engine.NewServer("/entry")
	go server.Listen()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
