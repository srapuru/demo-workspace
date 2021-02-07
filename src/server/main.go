
package main

import (
	"context"
	"log"
	"net"
	"fmt"
	"os"
	"generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":9999"
)

// server is used to implement  HelloService
type server struct { }




func main() {

	fmt.Printf("Please see the logfile server.log ")
    file, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    log.SetOutput(file)

	log.Print("==== Starting the server ")


	//listen on tcp port 
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//create server 
	srv := grpc.NewServer()
	generated.RegisterHelloServiceServer(srv, &server{})
	reflection.Register(srv)

	//start server 
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}


}


func (s *server) Hello(ctx context.Context, in *generated.HelloRequest) (*generated.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &generated.HelloResponse{Message: "Hello " + in.GetName()}, nil
}
