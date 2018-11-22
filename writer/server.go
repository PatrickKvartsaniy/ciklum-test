package main

import (
	"fmt"
	"io"
	"log"

	"github.com/PatrickKvartsaniy/ciklum-test/api"
)

// Server is used to implement customer.CustomerServer.
type Server struct{}

// NewServer - Server factory
func NewServer() *Server {
	return &Server{}
}

// CreateCustomer creates a new Customer
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
			CheckErr(err)
			return err
		}
		// create & save received customer to db
		customer := CreateCustomer(inCustomer)
		// check if customer does't exist - save
		if db.Find(&customer).RecordNotFound() {
			db.Save(&customer)
			out := fmt.Sprintf("User :%v  has been successfully added", customer.Name)
			log.Println(out)
		} else {
			updatedCustomer := UpdateCustomer(customer, inCustomer)
			db.Save(&updatedCustomer)
			out := fmt.Sprintf("User :%v  has been successfully updated", customer.Name)
			log.Println(out)
		}
	}
	return inStream.SendAndClose(&api.CustomerResponse{})
}
