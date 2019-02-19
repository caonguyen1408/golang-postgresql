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
	//insertProduct(pg_db)
	//deleteProduct(pg_db)
	//updateProduct(pg_db)
	getProduct(pg_db)
}

//get by id
func getProduct(dbRef *pg.DB) {
	newPI := &db.ProductItem{
		ID:   1,
		Name: "Product 2",
	}
	newPI.GetbyID(dbRef)
}

//update data in table
func updateProduct(dbRef *pg.DB) {
	newPI := &db.ProductItem{
		ID:   11,
		Name: "Product 6",
	}
	newPI.Update(dbRef)
}

//delete data in table
func deleteProduct(dbRef *pg.DB) {
	newPI := &db.ProductItem{
		Name: "Product 2",
	}
	newPI.Delete(dbRef)
}

//insert to table
func insertProduct(dbRef *pg.DB) {
	for i := 4; i < 7; i++ {
		newPI := &db.ProductItem{
			Name:  "Product " + string(i) + " ",
			Desc:  "Product" + string(i) + " desc",
			Image: "image pathaaaa",
			Price: 3.4,
			Features: struct {
				Name string
				Desc string
				Imp  int
			}{
				Name: "F" + string(i) + "",
				Desc: "F" + string(i) + " desc",
				Imp:  300,
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			IsActive:  true,
		}
		newPI.Save(dbRef)
	}
}
