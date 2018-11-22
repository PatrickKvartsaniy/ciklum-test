package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"

	"github.com/PatrickKvartsaniy/ciklum-test/api"

	"google.golang.org/grpc"
)

// gRPC server adress
var (
	gHost = os.Getenv("gHOST")
	gPort = os.Getenv("gPORT")
	gRPC  = gHost + ":" + gPort
)

// StreamCSV is implementation of  gRPC messages streaming
func StreamCSV(file multipart.File) (string, error) {
	var result string
	// connecting to gRPC server
	conn, err := grpc.Dial(gRPC, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer conn.Close()

	client := api.NewWriterClient(conn)

	ctx := context.Background()
	stream, err := client.CreateCustomer(ctx)

	reader := csv.NewReader(bufio.NewReader(file))
	log.Println("Start streaming")
	for {
		line, readerErr := reader.Read()
		// if file ends
		if readerErr == io.EOF {
			break
		} else if readerErr != nil {
			log.Fatal(err)
			return "", err
		}
		// skip csv head
		if line[0] == "id" {
			continue
		}
		// create & send gRPC customer
		customer := newCustomer(line)
		if err := stream.Send(customer); err != nil {
			log.Fatal(err)
		}
	}
	// graceful shutdown
	_, err = stream.CloseAndRecv()
	if err == io.EOF {
		log.Println("Stream has been closed")
		result = "Customers have been successufly added"
	} else if err != nil {
		result = fmt.Sprintf("Stream.CloseAndRecv() got error %v, want %v", err, nil)
		return result, err
	}
	return result, nil
}

// newCustomer is gRPC Customer  factory
func newCustomer(line []string) *api.Customer {
	return &api.Customer{
		Name:  line[1],
		Email: line[2],
		Phone: line[3],
	}
}
