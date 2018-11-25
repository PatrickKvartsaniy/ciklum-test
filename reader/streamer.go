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
		msg := fmt.Sprintf("An error occurred while connecting to grpc server. Error: %v", err)
		log.Fatal(msg)
		return msg, err
	}
	defer conn.Close()

	client := api.NewWriterClient(conn)

	ctx := context.Background()
	stream, err := client.CreateCustomer(ctx)

	reader := csv.NewReader(bufio.NewReader(file))
	log.Println("Start streaming csv file")
	for {
		line, readerErr := reader.Read()
		// if file is over
		if readerErr == io.EOF {
			break
		} else if readerErr != nil {
			msg := fmt.Sprintf("Something went wrong while trying read %v line from csv file. Error: %v", line[0], readerErr)
			log.Fatalf(msg)
			return msg, err
		}
		// skip csv head
		if line[0] == "id" {
			continue
		}
		// create & send gRPC customer
		customer := newCustomer(line)
		if err := stream.Send(customer); err != nil {
			log.Fatalf("An error occurred while trying to send customer %v data. Error: %v", line[0], err)
		}
	}
	// graceful shutdown
	_, err = stream.CloseAndRecv() //  _ - because we don't exprect any data from server
	if err == io.EOF {
		log.Println("Stream has been closed")
		result = "Customers have been successufly added"
	} else if err != nil {
		result = fmt.Sprintf("Stream.CloseAndRecv() got error %v, want %v", err, nil)
		return result, err
	}
	return result, nil
}

func newCustomer(line []string) *api.Customer {
	return &api.Customer{
		Name:  line[1],
		Email: line[2],
		Phone: line[3],
	}
}
