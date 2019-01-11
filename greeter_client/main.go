package main

import (
	"log"
	"os"
	"time"

	pb "grpc-helloworld/helloworld"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	startTime := time.Now()
	log.Printf("Start Time: %s", time.Now().String())
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	duration := time.Now().Sub(startTime)
	log.Printf("End Time: %s", time.Now().String())
	//log.Printf("Use Time: %d", duration)
	log.Printf("Use Time: %9.7f ms", float64(float64(duration)/1000/1000))
	log.Printf("Greeting: %s", r.Message)
}
