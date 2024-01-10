package main

import (
	"context"
	"log"

	pbproduct "github.com/suryatresna/samplegrpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcServerTarget = "localhost:50051"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.Dial(grpcServerTarget, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pbproduct.NewProductServiceClient(conn)

	for i := 0; i < 10; i++ {
		resp, err := client.GetProduct(ctx, &pbproduct.ProductRequest{})
		if err != nil {
			log.Fatalf("fail to call: %v", err)
		}
		log.Printf("response: %v", resp)
	}
}
