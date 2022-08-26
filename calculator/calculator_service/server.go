package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/VJ-Vijay77/gRPC/calculator/calculatorpb"

	"google.golang.org/grpc"
)

type server struct{

}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with %v \n",req)
	first := req.GetFirstNumber()
	second := req.GetSecondNumber()
	sum := first + second

	res := calculatorpb.SumResponse{
		Sum: sum,
	}

	return &res,nil

}

func main() {
	fmt.Println("Starting")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}

}
