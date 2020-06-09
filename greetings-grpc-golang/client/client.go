package main

import (
	"fmt"
	"context"
	"log"
	"google.golang.org/grpc"
	"github.com/fredoliveira-ca/greetings-grpc-java-golang/greetings-grpc-golang/greetpb"
)

func main() {

	fmt.Println("Hello! How are you?")
	javaTarget := "localhost:50052"

	connection, err := grpc.Dial(javaTarget, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer connection.Close()

	c := greetpb.NewGreetServiceClient(connection)

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Fred",
			SecondName:  "Oliveira",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Response)
}