package tools

import (
	"ciklum/writer/config"
	"ciklum/writer/models"

	"github.com/jinzhu/gorm"
)

func CreateEngine() *gorm.DB {
	db, err := gorm.Open("postgres", config.DbRoute)
	CheckErr(err)
	return db
}

func MakeMigrations() {
	db := CreateEngine()
	defer db.Close()

	db.AutoMigrate(&models.Customer{})
}
