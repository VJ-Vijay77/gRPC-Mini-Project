package main

import (
	"context"
	"fmt"
	"log"

	"github.com/VJ-Vijay77/gRPC/greet/greetpb"
	"google.golang.org/grpc"
)


func main() {
	fmt.Println("Hello im a client!")

	cc,err := grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err != nil {
		log.Fatalln("could not connect ",err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	// fmt.Printf("Created Client:%f",c)
	doUnary(c)
	
}


func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Vijay",
			LastName: "Dinesh",
		},
	}

	res,err := c.Greet(context.Background(),req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC : %v", err)
	}
	log.Printf("Response from Greet : %v",res.Result)
}