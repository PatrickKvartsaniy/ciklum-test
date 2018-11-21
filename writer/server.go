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
		if err := db.Save(&customer).Error; err != nil {
			log.Fatal(err)
			//  If it exists - update fileds
			updatedCustomer := UpdateCustomer(customer, inCustomer)
			db.Save(&updatedCustomer)
			out := &api.CustomerResponse{
				Response: fmt.Sprintf("User :%v  has been successfully updated", customer.Name),
			}
			inStream.Send(out)
		} else {
			out := &api.CustomerResponse{
				Response: fmt.Sprintf("User :%v  has been successfully added", customer.Name),
			}
			inStream.Send(out)
		}
	}
	return nil
}
