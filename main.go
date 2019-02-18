package main

import (
	"log"

	db "./db"
)

func main() {
	log.Printf("hello world...!!!.\n")
	db.Connect()
}
