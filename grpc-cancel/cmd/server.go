package main

import (
	"fmt"
	pb "github.com/okharch/golang-grpc-examples/grpc-cancel/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {
	pb.UnimplementedTimeServiceServer
}

func (s *server) StreamTime(_ *pb.StreamTimeRequest, stream pb.TimeService_StreamTimeServer) error {
	for {
		select {
		case <-stream.Context().Done():
			fmt.Println("Client has cancelled the context")
			return nil
		default:
			timeStr := time.Now().Format(time.RFC3339)
			if err := stream.Send(&pb.StreamTimeResponse{Time: timeStr}); err != nil {
				log.Printf("error on server sending stream response: %s", err)
				return err
			}
			time.Sleep(time.Second)
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTimeServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
