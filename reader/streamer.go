package main

import (
	"bufio"
	"ciklum/api"
	"ciklum/reader/tools"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	"google.golang.org/grpc"
)

// StreamCSV is implementation of  gRPC messages streaming
func StreamCSV(file multipart.File) {
	reader := csv.NewReader(bufio.NewReader(file))
	// connecting to gRPC server
	conn, err := grpc.Dial("127.0.0.1:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cant connect to grpc. Pls check if port is correct ")
	}
	defer conn.Close()

	client := api.NewWriterClient(conn)

	ctx := context.Background()
	stream, err := client.CreateCustomer(ctx)
	tools.CheckErr(err)
	log.Println("Start streaming")
	for {
		line, err := reader.Read()
		//  checking if stream is over
		if err == io.EOF {
			log.Println("Stream has been closed")
			break
		}
		if err != nil {
			tools.CheckErr(err)
			break
		}
		//  skip csv head
		if line[0] == "id" {
			continue
		}
		// create & send gRPC customer
		customer := newCustomer(line)
		if err := stream.Send(customer); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, customer, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	tools.CheckErr(err)
	fmt.Println(reply)
}

// newCustomer is gRPC customer factory
func newCustomer(line []string) *api.Customer {
	return &api.Customer{
		Name:  line[1],
		Email: line[2],
		Phone: line[3],
	}
}
