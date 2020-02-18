package main

import (
	"./engine"
)

func main() {
	server := engine.NewServer("/entry")
	go server.Listen()
}
