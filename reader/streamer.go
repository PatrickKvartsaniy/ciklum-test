package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"io"
	"log"
	"mime/multipart"
	"os"

	"ciklum-test/reader/tools"

	"github.com/PatrickKvartsaniy/ciklum-test/api"

	"google.golang.org/grpc"
)

// gRPC server adress
var (
	gHost = os.Getenv("gHost")
	gPort = os.Getenv("gPort")
	gRPC  = gHost + ":" + gPort
)

// StreamCSV is implementation of  gRPC messages streaming
func StreamCSV(file multipart.File) {

	// connecting to gRPC server
	conn, err := grpc.Dial(gRPC, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Cant connect to grpc. Pls check if port is correct ")
	}
	defer conn.Close()

	client := api.NewWriterClient(conn)

	ctx := context.Background()
	stream, err := client.CreateCustomer(ctx)
	tools.CheckErr(err)

	reader := csv.NewReader(bufio.NewReader(file))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			tools.CheckErr(err)
		}
		// skip csv head
		if line[0] == "id" {
			continue
		}
		// create & send gRPC customer
		customer := newCustomer(line)
		if err := stream.Send(customer); err != nil {
			tools.CheckErr(err)
		}
	}
	// graceful shutdown
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	} else {
		log.Println("Customers have been successfully saved")
	}
}

// newCustomer is gRPC Customer  factory
func newCustomer(line []string) *api.Customer {
	return &api.Customer{
		Name:  line[1],
		Email: line[2],
		Phone: line[3],
	}
}
