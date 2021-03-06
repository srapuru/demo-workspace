package main

import (
	"context"
	"generated"
	"log"
	"os"
	"fmt"
	"time"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:9999"
	defaultName = "GRPC"
)

func main() {

	fmt.Printf("Please see the logfile client.log \n")
    file, err := os.OpenFile("client.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    log.SetOutput(file)

	log.Print("==== Starting the client ")

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := generated.NewHelloServiceClient(conn)

	// Contact the server and print out its response.
	name := defaultName

	log.Print("==== Client started successfully ");

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()
	res, err := c.Hello(ctx, &generated.HelloRequest{Name: name})
		if err != nil {
				log.Fatalf("ERROR in HelloRequest : %v", err)
		}
	log.Printf("Response from the HelloRequest : %s", res.GetMessage())
}
