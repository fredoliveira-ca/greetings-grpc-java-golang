package com.fredoliveira.greetingsgrpcjava.server;

import io.grpc.greet.GreetRequest;
import io.grpc.greet.GreetResponse;
import io.grpc.greet.GreetServiceGrpc;
import io.grpc.stub.StreamObserver;

public class GreetServiceImpl extends GreetServiceGrpc.GreetServiceImplBase {

    @Override
    public void greet(GreetRequest request, StreamObserver<GreetResponse> responseObserver) {
        String greeting = "Greet, " +
                request.getGreeting().getFirstName() +
                " " +
                request.getGreeting().getSecondName();

        GreetResponse response = GreetResponse.newBuilder()
                .setResponse(greeting)
                .build();

        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }
}
