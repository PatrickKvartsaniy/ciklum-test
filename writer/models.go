package main

/*
	I know that gORM isn't good for production. I just used it here for save my time
*/

import (
	"log"

	"github.com/PatrickKvartsaniy/ciklum-test/api"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateEngine() *gorm.DB {
	uri = SetupDB()
	db, err := gorm.Open("postgres", uri)
	if err != nil {
		log.Fatalf("Something went wrong while connecting to database %v. Error: %v.", db, err)
	}
	db.LogMode(false) // for clear logs
	return db
}

// Customer gorm model
type Customer struct {
	gorm.Model
	Name  string `gorm:"not null;"`
	Email string `gorm:"not null;unique"`
	Phone string `gorm:"not null;unique"`
}

func CreateCustomer(in *api.Customer) *Customer {
	return &Customer{
		Name:  in.Name,
		Email: in.Email,
		Phone: "(+44)" + in.Phone,
	}
}

func UpdateCustomer(exist Customer, in *Customer) *Customer {
	exist.Name = in.Name
	exist.Email = in.Email
	exist.Phone = in.Phone
	return &exist
}

func MakeMigrations() {
	db := CreateEngine()
	defer db.Close()

	db.AutoMigrate(&Customer{})
}
