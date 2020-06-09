package com.fredoliveira.greetingsgrpcjava.server;

import java.io.IOException;

import io.grpc.Server;
import io.grpc.ServerBuilder;

public class GrpcServer {
    public static void main(String[] args) throws IOException, InterruptedException {
        Server server = ServerBuilder
                .forPort(50052)
                .addService(new GreetServiceImpl()).build();

        server.start();
        server.awaitTermination();
    }
}
