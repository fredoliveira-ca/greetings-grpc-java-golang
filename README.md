# greetings-grpc-java-golang

- Go server port: 50051
``` bash
go run server/server.go
```

- Java server port: 50052
-- Required Java 14
``` bash
java -jar com.fredoliveira.greetingsgrpcjava.server.GrpcServer.jar
```


- Go client (You have to run java server)
``` bash
go run client/client.go 
```

- Go client (You have to run java server)
-- Required Java 14
``` bash
java -jar com.fredoliveira.greetingsgrpcjava.server.GrpcClient.jar
```