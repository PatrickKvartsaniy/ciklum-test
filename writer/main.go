package main

import (
	"ciklum/api"
	"ciklum/writer/tools"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db := tools.CreateEngine()
	defer db.Close()
	tools.MakeMigrations()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	api.RegisterWriterServer(s, &server{})
	s.Serve(lis)
	// db.Create(&customer)
	// me := models.Customer{Name: "Petro", Email: "kvartsaniy@gmai.com", Phone: "+380938471506"}
	// fmt.Println(me.Name)
}
