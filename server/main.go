package main

import (
	"context"
	"grpc-go/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {}

func (s *Server) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Msg: "Hello " + request.GetName()}, nil
}


func main(){
	listen, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}