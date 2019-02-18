package main

import (
	"log"
	"time"

	"github.com/go-pg/pg"

	db "./db"
)

func main() {
	log.Printf("hello world...!!!.\n")
	pg_db := db.Connect()
	insertProduct(pg_db)
}
func insertProduct(dbRef *pg.DB) {
	newPI := &db.ProductItem{
		Name:  "Product 2",
		Desc:  "Product 2 desc",
		Image: "image pathaaaa",
		Price: 3.4,
		Features: struct {
			Name string
			Desc string
			Imp  int
		}{
			Name: "F2",
			Desc: "F2 desc",
			Imp:  300,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  true,
	}
	newPI.Save(dbRef)
}
