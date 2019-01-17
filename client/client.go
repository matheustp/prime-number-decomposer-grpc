package main

import (
	"context"
	"io"
	"log"

	pndpb "github.com/matheustp/prime-number-decomposer-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Panicf("Error when dialing: %v", err)
	}
	defer cc.Close()
	c := pndpb.NewPrimeNumberDecomposerServiceClient(cc)
	req := &pndpb.PrimeNumberDecomposerRequest{
		Number: 42,
	}
	stream, err := c.PrimeNumberDecompose(context.Background(), req)
	if err != nil {
		log.Panicf("Error when calling function: %v", err)
	}
	result := []int32{}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panicf("Error fetching message: %v", err)
		}
		result = append(result, msg.GetResult())
	}
	log.Printf("The prime number decomposers for %v are: %v", req.GetNumber(), result)
}
