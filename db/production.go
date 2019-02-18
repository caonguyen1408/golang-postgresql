package db

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type ProductItem struct {
	RefPointer int      `sql:"-"`
	tableName  struct{} `sql:"product_items_collection"`
	ID         int      `sql:"id,pk"`
	Name       string   `sql:"name,unique"`
	Desc       string   `sql:"desc"`
	Image      string   `sql:"image"`
	Price      float64  `sql:"price,type:real"`
	Features   struct {
		Name string
		Desc string
		Imp  int
	} `sql:"feature,type:jsonb"`
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
	IsActive  bool      `sql:"is_active"`
}

// Create table
func createProdItemtable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		//check if not exists create new
		IfNotExists: true,
	}
	createErr := db.CreateTable(&ProductItem{}, opts)
	if createErr != nil {
		log.Printf("Error while create table, Reason %v\n", createErr)
		return createErr
	}
	log.Printf("Table productItem created successful. \n")
	return nil
}

// Insert in to table
// define shortcut name (pi *ProductItem)
func (pi *ProductItem) Save(db *pg.DB) error {
	insertErr := db.Insert(pi)
	if insertErr != nil {
		log.Printf("Error while inserting new item into DB, Reason: %v\n", insertErr)
		return insertErr
	}
	log.Printf("ProductItem %s insert successfull\n ", pi.Name)
	return nil
}
