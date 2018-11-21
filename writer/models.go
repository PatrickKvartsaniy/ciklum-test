package main

import (
	"ciklum-test/api"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Customer db model
type Customer struct {
	gorm.Model
	Name  string `gorm:"not null;"`
	Email string `gorm:"not null;unique"`
	Phone string `gorm:"not null;unique"`
}

// CreateCustomerModel factory
func CreateCustomer(in *api.Customer) *Customer {
	return &Customer{
		Name:  in.Name,
		Email: in.Email,
		Phone: in.Phone,
	}
}

// UpdateCustomer will update customer in db
func UpdateCustomer(exist *Customer, in *api.Customer) *Customer {
	exist.Name = in.Name
	exist.Email = in.Email
	exist.Phone = in.Phone
	return exist
}
