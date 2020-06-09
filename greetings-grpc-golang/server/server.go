package main

import (
	"fmt"
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/fredoliveira-ca/greetings-grpc-java-golang/greetings-grpc-golang/greetpb"
)

type server struct{}

func main() {
	fmt.Println("Hello! Ready to serve and listen you...")

	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	newServer := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(newServer, &server{})

	if err := newServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	secondName := req.GetGreeting().GetSecondName()
	response := "Hello " + firstName + " " + secondName + "! I'm great!"
	res := &greetpb.GreetResponse{
		Response: response,
	}
	return res, nil
}