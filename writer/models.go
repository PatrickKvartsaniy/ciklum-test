package main

import (
	"github.com/PatrickKvartsaniy/ciklum-test/api"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

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
		Phone: in.Phone,
	}
}

func UpdateCustomer(exist *Customer, in *api.Customer) *Customer {
	exist.Name = in.Name
	exist.Email = in.Email
	exist.Phone = in.Phone
	return exist
}
