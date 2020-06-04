package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/frybin/laforge/grpc-alpha/laforge_proto"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedLaforgeServer
}

/* TEST MESSAGES */

//Ping Info
func (s *server) GetPing(ctx context.Context, in *pb.PingRequest) (*pb.PingReply, error) {
	log.Printf("Received: %v | ID: %v", in.GetName(), in.GetId())
	return &pb.PingReply{Name: "Hello " + in.GetName(), Id: in.GetId()}, nil
}

//HostTest Info
func (s *server) GetHostTest(ctx context.Context, in *pb.HostTestRequest) (*pb.HostTestReply, error) {
	log.Printf("Got Host: %v | ID: %v | IP: %s | OS: %s", in.GetName(), in.GetId(), in.GetIp(), in.GetOs())
	return &pb.HostTestReply{Name: in.GetName(), Id: in.GetId(), Ip: in.GetIp(), Os: in.GetOs()}, nil
}


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLaforgeServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}