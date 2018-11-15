package tools

import (
	"ciklum/reader/configs"
	"ciklum/reader/models"

	"github.com/jinzhu/gorm"
)

func CreateEngine() *gorm.DB {
	db, err := gorm.Open("postgres", configs.DbRoute)
	CheckErr(err)
	return db
}

func MakeMigrations() {
	db := CreateEngine()
	defer db.Close()

	db.AutoMigrate(&models.Customer{})
}
