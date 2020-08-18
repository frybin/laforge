package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Command struct {
	gorm.Model
	Name  string
	Price uint
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Command{})

	// Create
	db.Create(&Command{Name: "L1212", Price: 1000})

	// Read
	var command Command
	db.First(&command, 1)                   // find product with id 1
	db.First(&command, "name = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&command).Update("Price", 2000)

	// Delete - delete product
	// db.Delete(&command)
}
