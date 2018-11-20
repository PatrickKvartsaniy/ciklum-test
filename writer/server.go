package main

import (
	"io"
	"log"

	"github.com/PatrickKvartsaniy/ciklum-test/api"

	"github.com/PatrickKvartsaniy/ciklum-test/writer/models"
	"github.com/PatrickKvartsaniy/ciklum-test/writer/tools"
)

// Server is used to implement customer.CustomerServer.
type Server struct{}

// CreateCustomer creates a new Customer
func (s *Server) CreateCustomer(inStream api.Writer_CreateCustomerServer) error {

	db := tools.CreateEngine()
	defer db.Close()

	for {
		inCustomer, err := inStream.Recv()
		//  checking if stream is over
		if err == io.EOF {
			log.Println("Stream has been closed")
			return nil
		}
		if err != nil {
			tools.CheckErr(err)
			return err
		}
		// create & save received customer to db
		customer := models.CreateCustomerModel(inCustomer)
		if err := db.Save(&customer).Error; err != nil {
			updatedCustomer := models.UpdateCustomer(customer, inCustomer)
			db.Save(&updatedCustomer)
			log.Printf("User :%v  has been successfully updated", customer.Name)
		}
		log.Printf("User :%v  has been successfully added", customer.Name)
	}
}

// NewServer - Server factory
func NewServer() *Server {
	return &Server{}
}
