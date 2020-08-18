package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	dbDSN = "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"
)

// Task to be assigned
type Task struct {
	ClientID  string `gorm:"primary_key"`
	TaskID    int32  `gorm:"primary_key;auto_increment:false"`
	CommandID int32
	Args      string
	Completed bool `gorm:"default:false"`
}

// TempURL to be assigned
type TempURL struct {
	URLPost  string `gorm:"unique;not null"`
	FilePath string `gorm:"not null"`
	Valid    bool   `gorm:"default:true"`
	gorm.Model
}

// OpenDB function used to connect to dataase and returns db object
func OpenDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	// db, err := gorm.Open(postgres.Open(dbDSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

// CreateDB function that auto migrates and creates database if it does not exist
func CreateDB(db *gorm.DB) {
	db.AutoMigrate(&Task{}, &TempURL{})
}

func test() {

	db := OpenDB()
	CreateDB(db)
	db.Create(&Task{ClientID: "Test", TaskID: 2, CommandID: 8, Args: "Bob", Completed: true})
	tasks := make([]Task, 0)
	db.Find(&tasks)
	fmt.Println(tasks)
}
