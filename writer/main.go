package main

import (
	"flag"
	"log"
	"net"

	"github.com/PatrickKvartsaniy/ciklum-test/api"
	"github.com/PatrickKvartsaniy/ciklum-test/writer/tools"

	"google.golang.org/grpc"
)

func main() {
	//  Auto migrations for models
	migrations := flag.Bool("makemigrations", false, "run migrations")
	port := flag.String("port", "5001", "port to run")
	flag.Parse()

	if *migrations {
		tools.MakeMigrations()
	}
	lis, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	api.RegisterWriterServer(s, NewServer())
	log.Println("Server running on 127.0.0.1:" + *port)
	s.Serve(lis)
}
