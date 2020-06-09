package com.fredoliveira.greetingsgrpcjava.client;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.greet.GreetRequest;
import io.grpc.greet.GreetResponse;
import io.grpc.greet.GreetServiceGrpc;
import io.grpc.greet.Greeting;

public class GrpcClient {
    public static void main(String[] args) {
        int golangPort = 50051;

        ManagedChannel channel = ManagedChannelBuilder.forAddress("localhost", golangPort)
                .usePlaintext()
                .build();

        GreetServiceGrpc.GreetServiceBlockingStub stub
                = GreetServiceGrpc.newBlockingStub(channel);

        GreetResponse response = stub.greet(
                GreetRequest.newBuilder()
                    .setGreeting(Greeting.newBuilder()
                            .setFirstName("Fred")
                            .setSecondName("Oliveira")
                            .build()).build()
        );

        System.out.println(response.toString());

        channel.shutdown();
    }
}

