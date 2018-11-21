package main

import (
	"log"

	"github.com/jinzhu/gorm"
)

func CreateEngine() *gorm.DB {
	db, err := gorm.Open("postgres", DbRoute)
	CheckErr(err)
	return db
}

func MakeMigrations() {
	db := CreateEngine()
	defer db.Close()

	db.AutoMigrate(&Customer{})
}

func CheckErr(err error) {
	if err != nil {
		// panic(err)
		log.Println(err.Error())
	}
}
