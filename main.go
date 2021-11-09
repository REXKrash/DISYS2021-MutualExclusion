package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"log"
	"net"
	"os"

	pb "token-program/routeguide"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTokenServiceServer
}

func (s *server) FindLeaderRequest(ctx context.Context, request *pb.LeaderRequest) (*pb.LeaderResponse, error) {
	fmt.Println("Received LeaderRequest from", targetPort, "looking for leader:", request.LookingForLeader, "lowest port:", request.LowestPort)
	lookingForLeader = request.LookingForLeader

	if request.LowestPort <= lowestPort {
		lowestPort = request.LowestPort
		if lowestPort == listenPort {
			lookingForLeader = false
			hasToken = true
			fmt.Println("Found leader, it is me!")
		}
	}
	return &pb.LeaderResponse{LookingForLeader: lookingForLeader, LowestPort: lowestPort}, nil
}

func (s *server) SendTokenRequest(ctx context.Context, request *pb.TokenRequest) (*pb.TokenResponse, error) {
	lookingForLeader = false
	hasToken = true
	fmt.Println("Received TokenRequest")
	return &pb.TokenResponse{Status: 1, Message: "Response"}, nil
}

var targetPort int32
var listenPort int32
var lowestPort int32
var lookingForLeader = true
var hasToken = false

func main() {
	fmt.Println("Welcome to FijiFunCoin, safely distributed unhackable coin system")
	if len(os.Args) > 3 {
		fmt.Println("Y U no give ports!")
		if len(os.Args) > 2 {
			fmt.Println("TargetPort: Missing")
		}
		if len(os.Args) > 1 {
			fmt.Println("ListenPort: Missing")
		}
		return
	}
	_targetPort, err1 := strconv.Atoi(os.Args[2])
	if err1 != nil {
		log.Fatalf("Bad Port")
	}
	_listenPort, err2 := strconv.Atoi(os.Args[1])
	if err2 != nil {
		log.Fatalf("Bad Port")
	}
	targetPort = int32(_targetPort)
	listenPort = int32(_listenPort)
	lowestPort = listenPort

	time.Sleep(10 * time.Second)
	go runServer()
	time.Sleep(1 * time.Second)
	go runClient()
	for {
		time.Sleep(2 * time.Second)
	}
}

func runClient() {
	address := "localhost:" + strconv.Itoa(int(targetPort))
	connection, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer connection.Close()
	client := pb.NewTokenServiceClient(connection)
	ctx := context.Background()

	for {
		if lookingForLeader {
			fmt.Println("Sending LeaderRequest to", targetPort)
			client.FindLeaderRequest(ctx, &pb.LeaderRequest{LookingForLeader: lookingForLeader, LowestPort: lowestPort})
		} else {
			break
		}
		time.Sleep(5 * time.Second)
	}
	for {
		if hasToken {
			fmt.Println("Sending TokenRequest to", targetPort)
			_, err = client.SendTokenRequest(ctx, &pb.TokenRequest{Status: 1, Message: "Request!"})
			if err != nil {
				log.Fatalf("Failed to send TokenRequest: %v", err)
			}
			hasToken = false
		}
		time.Sleep(5 * time.Second)
	}
}

func runServer() {
	fmt.Println("--- SERVER APP ---")

	address := "localhost:" + strconv.Itoa(int(listenPort))
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterTokenServiceServer(s, &server{})

	fmt.Println("Listen port:", listenPort)
	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
