package main

import (
	"github.com/PatrickKvartsaniy/ciklum-test/api"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Customer model struct
type Customer struct {
	gorm.Model
	Name  string `gorm:"not null;"`
	Email string `gorm:"not null;unique"`
	Phone string `gorm:"not null;unique"`
}

// CreateCustomer model
func CreateCustomer(in *api.Customer) *Customer {
	return &Customer{
		Name:  in.Name,
		Email: in.Email,
		Phone: "(+44)" + in.Phone,
	}
}

// UpdateCustomer model rows
func UpdateCustomer(exist Customer, in *Customer) *Customer {
	exist.Name = in.Name
	exist.Email = in.Email
	exist.Phone = in.Phone
	return &exist
}
