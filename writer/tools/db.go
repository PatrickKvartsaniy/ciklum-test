package tools

import (
	"ciklum/writer/configs"
	"ciklum/writer/models"

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
