package main

import (
	"log"
)

func main() {
	app := App{8000}
	log.Printf("Hello, AI")
	app.Start()
}
