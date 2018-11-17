package main

import (
	"bufio"
	"ciklum/api"
	"ciklum/reader/tools"
	"context"
	"encoding/csv"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func NewCustomer(line []string) *api.Customer {
	return &api.Customer{
		Name:  line[1],
		Email: line[2],
		Phone: line[3],
	}
}

func createCustomer(client api.WriterClient, customer *api.Customer) {
	resp, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatalf("Could not create Customer: %v", err)
	}
	if resp.Success {
		log.Printf("A new Customer has been added")
	}
}

func main() {

	file, err := os.Open("data.csv")
	tools.CheckErr(err)
	defer file.Close()
	reader := csv.NewReader(bufio.NewReader(file))

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	tools.CheckErr(err)
	defer conn.Close()

	client := api.NewWriterClient(conn)

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		customer := NewCustomer(line)
		createCustomer(client, customer)
	}
}
