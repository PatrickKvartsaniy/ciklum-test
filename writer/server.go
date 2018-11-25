package main

import (
	"fmt"
	"io"
	"log"

	"github.com/PatrickKvartsaniy/ciklum-test/api"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) CreateCustomer(inStream api.Writer_CreateCustomerServer) error {
	log.Println("Start streaming")
	db := CreateEngine()
	defer db.Close()

	for {
		inCustomer, err := inStream.Recv()
		//  checking if stream is over
		if err == io.EOF {
			log.Println("Stream has been closed")
			return nil
		}
		if err != nil {
			log.Fatal(err)
			return err
		}

		// create received customer object
		customer := CreateCustomer(inCustomer)
		// check if customer does't exist - save
		var a Customer
		if db.Debug().First(&a, "email = ? OR phone = ?", customer.Email, customer.Phone).RecordNotFound() {
			db.Save(&customer)
			out := fmt.Sprintf("User :%v  has been successfully created", customer.Name)
			log.Println(out)
		} else {
			// else - update fields
			var existCustomer Customer
			db.Where("email = ?", customer.Email).Find(&existCustomer)
			updatedCustomer := UpdateCustomer(existCustomer, customer)
			db.Save(&updatedCustomer)
			out := fmt.Sprintf("User :%v  has been successfully updated", customer.Name)
			log.Println(out)
		}
	}
	return inStream.SendAndClose(&api.CustomerResponse{})
}
