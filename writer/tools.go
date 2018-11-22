package main

import (
	"log"

	"github.com/jinzhu/gorm"
)

func CreateEngine() *gorm.DB {
	db, err := gorm.Open("postgres", DbRoute)
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(false)
	return db
}

func MakeMigrations() {
	db := CreateEngine()
	defer db.Close()

	db.AutoMigrate(&Customer{})
}
