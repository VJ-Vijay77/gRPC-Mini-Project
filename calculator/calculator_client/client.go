package main

import (
	"context"
	"fmt"
	"log"
	"github.com/VJ-Vijay77/gRPC/calculator/calculatorpb"
	"google.golang.org/grpc"
)


func main() {
	fmt.Println("Hello im a client!")

	cc,err := grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err != nil {
		log.Fatalln("could not connect ",err)
	}
	defer cc.Close()
	c := calculatorpb.NewSumServiceClient(cc)
	// fmt.Printf("Created Client:%f",c)
	doUnary(c)
	
}


func doUnary(c calculatorpb.SumServiceClient) {
	fmt.Println("Starting to do a Calculator Unary RPC...")
	req := &calculatorpb.SumRequest{
		FirstNumber: 45,
		SecondNumber: 5,
	}

	res,err := c.Sum(context.Background(),req)
	if err != nil {
		log.Fatalf("error while calling SUM RPC : %v", err)
	}
	log.Printf("Response from SUM : %v",res.Sum)
}