package tools

import (
	"github.com/PatrickKvartsaniy/ciklum-test/writer/config"
	"github.com/PatrickKvartsaniy/ciklum-test/writer/models"

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
