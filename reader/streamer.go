package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"io"
	"log"
	"mime/multipart"
	"os"
	"sync"

	"ciklum-test/reader/tools"

	"ciklum-test/api"

	"google.golang.org/grpc"
)

// gRPC server adress
var gRPC = os.Getenv("gRPC")

// StreamCSV is implementation of  gRPC messages streaming
func StreamCSV(file multipart.File) {

	// connecting to gRPC server
	conn, err := grpc.Dial(
		gRPC+":5001", //writer's docker address
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatalf("Cant connect to grpc. Pls check if port is correct ")
	}
	defer conn.Close()

	client := api.NewWriterClient(conn)

	ctx := context.Background()
	stream, err := client.CreateCustomer(ctx)
	tools.CheckErr(err)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	// Sending
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		reader := csv.NewReader(bufio.NewReader(file))
		for {
			line, err := reader.Read()
			if err == io.EOF {
				break
			}
			//  skip csv head
			if line[0] == "id" {
				continue
			}
			// create & send gRPC customer
			customer := newCustomer(line)
			stream.Send(customer)
		}
		// graceful shutdown
		stream.CloseSend()
	}(wg)

	// Receiving
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			response, err := stream.Recv()
			//  checking if stream is over
			if err == io.EOF {
				log.Println("Stream has been closed")
				return
			} else if err != nil {
				tools.CheckErr(err)
				return
			}
			log.Println(response)
		}
	}(wg)
	wg.Wait()
}

// newCustomer is gRPC Customer  factory
func newCustomer(line []string) *api.Customer {
	return &api.Customer{
		Name:  line[1],
		Email: line[2],
		Phone: line[3],
	}
}
