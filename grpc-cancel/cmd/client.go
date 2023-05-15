package main

import (
	"context"
	"fmt"
	pb "github.com/okharch/golang-grpc-examples/grpc-cancel/api"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTimeServiceClient(conn)

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	stream, err := client.StreamTime(ctx, &pb.StreamTimeRequest{})
	if err != nil {
		log.Fatalf("could not stream time: %v", err)
	}

	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				return
			}
			fmt.Println("Received time:", res.GetTime())
		}
	}()

	<-ctx.Done()
	fmt.Println("Context has been cancelled on the client side")
	log.Printf("client waits 10 sec before exit...")
	time.Sleep(time.Second * 10)
}
