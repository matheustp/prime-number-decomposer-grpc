package main

import (
	"log"
	"net"

	pndpb "github.com/matheustp/prime-number-decomposer-grpc/pb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) PrimeNumberDecompose(req *pndpb.PrimeNumberDecomposerRequest, stream pndpb.PrimeNumberDecomposerService_PrimeNumberDecomposeServer) error {
	k := int32(2)
	n := req.GetNumber()
	for n > 1 {
		if n%k == 0 {
			res := &pndpb.PrimeNumberDecomposerResponse{
				Result: k,
			}
			stream.Send(res)
			n = n / k
		} else {
			k++
		}
	}
	return nil
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Panicf("Error when trying to listen: %v", err)
	}
	defer l.Close()
	log.Println("Started to listen")
	s := grpc.NewServer()
	pndpb.RegisterPrimeNumberDecomposerServiceServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Panicf("Error when registering rpc: %v", err)
	}

}
