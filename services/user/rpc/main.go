package rpc

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"user/repositories"

	"google.golang.org/grpc"
	pb "protobuf/_go/user"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	userRepo := repositories.UserRepository{}
	users := userRepo.GetAll()
	b, _ := json.Marshal(users)
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: string(b)}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello2 " + in.GetName()}, nil
}

func RunRPCServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
