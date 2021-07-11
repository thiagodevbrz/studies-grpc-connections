package main

import (
	"log"

	"net"

	"github.com/thiagodevbrz/grpc-exercise/pb"
	"github.com/thiagodevbrz/grpc-exercise/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:50001")

	if err != nil {
		log.Fatalf("Could not serve: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)

	serverError := grpcServer.Serve(lis)

	if serverError != nil {
		log.Fatalf("Could not serve: %v", serverError)
	}

	log.Println("Server listening at port 50001")

}
