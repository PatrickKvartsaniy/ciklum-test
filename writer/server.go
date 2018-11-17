package main

import (
	"ciklum/api"
	"fmt"

	"golang.org/x/net/context"
)

const (
	port = ":50051"
)

// server is used to implement customer.CustomerServer.
type server struct {
}

// CreateCustomer creates a new Customer
func (s *server) CreateCustomer(ctx context.Context, in *api.Customer) (*api.CustomerResponse, error) {
	fmt.Println(in)
	return &api.CustomerResponse{Success: true}, nil
}
