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

// Delete data in table
func (pi *ProductItem) Delete(db *pg.DB) error {
	_, deleteErr := db.Model(pi).Where("name= ?name").Delete()
	if deleteErr != nil {
		log.Printf("Error while inserting new item into DB, Reason: %v\n", deleteErr)
		return deleteErr
	}
	log.Printf("ProductItem %s delete successfull\n ", pi.Name)
	return nil
}

//update database
func (pi *ProductItem) Update(db *pg.DB) error {
	tx, txErr := db.Begin()
	if txErr != nil {
		log.Printf("Error while open tx, Reason %v\n", txErr)
		return txErr
	}
	_, updateErr := tx.Model(pi).Set("name = ?name").Where("id= ?id").Update()
	if updateErr != nil {
		log.Printf("Error while inserting new item into DB, Reason: %v\n", updateErr)
		tx.Rollback()
		return updateErr
	}
	tx.Commit()
	log.Printf("ProductItem %s delete successfull\n ", pi.Name)
	return nil
}

// get by id
func (pi *ProductItem) GetbyID(db *pg.DB) error {
	var pis []ProductItem
	getByID := db.Model(&pis).
		Where("id = ?0", pi.ID).
		WhereOr("name = ?0", pi.Name).
		Order("id desc").
		Select()
	if getByID != nil {
		log.Printf("Error while inserting new item into DB, Reason: %v\n", getByID)
		return getByID
	}
	log.Printf("ProductItem %s delete successfull\n ", pis)
	return nil
}
