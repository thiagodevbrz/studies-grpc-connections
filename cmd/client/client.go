package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/thiagodevbrz/grpc-exercise/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50001", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC Server at port 50001: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	// AddUser(client)
	// AddUserVerbose(client)
	// AddUsers(client)
	AddUserStreamBoth(client)

}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Beatriz Bitencourt",
		Email: "beatriz_bitencourt.silva@hotmail.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not receive the stream message: %v", err)
		}

		fmt.Println("Status:", stream.Status)
	}

}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Beatriz Bitencourt",
		Email: "beatriz_bitencourt.silva@hotmail.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUsers(client pb.UserServiceClient){
	reqs := []*pb.User{
		&pb.User{
			Id: "t1",
			Name: "Thiago P"
			Email: "thiago.dev.brz@gmail.com",
		},
		&pb.User{
			Id: "t2",
			Name: "Beatriz Bitencourt Silva"
			Email: "beatriz_bitencourt.silva@hotmail.com",
		},
		&pb.User{
			Id: "t3",
			Name: "Lívia Bitencourt Pereira"
			Email: "livia.dev.brz@gmail.com",
		}
	}

	stream,err:= client.AddUsers(context.Background())

	if err != nill {
		log.Fatalf("Error creating request stream: %v", err)
	}

	for _, req := range reqs {
		fmt.Println(req)

		stream.Send(req)

		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()

	if err != nill {
		log.Fatalf("Error receiving response stream: %v", err)
	}
	
	fmt.Println(res)

}

func AddUserStreamBoth(client pb.UserServiceClient){
	stream, err := client.AddUserStreamBoth(context.Background())

	if err != nill {
		log.Fatalf("Error creating request stream: %v", err)
	}

	reqs := []*pb.User{
		&pb.User{
			Id: "t1",
			Name: "Thiago P"
			Email: "thiago.dev.brz@gmail.com",
		},
		&pb.User{
			Id: "t2",
			Name: "Beatriz Bitencourt Silva"
			Email: "beatriz_bitencourt.silva@hotmail.com",
		},
		&pb.User{
			Id: "t3",
			Name: "Lívia Bitencourt Pereira"
			Email: "livia.dev.brz@gmail.com",
		}
	}

	wait := make(chan int)

	go func(){
		for _,req := range reqs {
			fmt.Println("Sending user :", req.Name)

			stream.Send((req))

			time.Sleep(time.Second * 2)
		}

		stream.CloseSend()
	}()

	go func(){
		for {

			res, err := stream.Recv()

			if err == io.EOF {
				fmt.Println("Finished fetching stream");
				break
			}

			if err != nill {
				log.Fatalf("Error receiving data stream: %v", err)
				break
			}

			fmt.Printf("Recebendo user %v com status %v". res.GetUser().getName(), res.GetStatus())
		}

		close(wait)
	}()

	<-wait



}
