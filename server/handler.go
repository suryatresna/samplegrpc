package main

import (
	"context"

	pbproduct "github.com/suryatresna/samplegrpc/proto"
)

type Handler struct {
	pbproduct.UnimplementedProductServiceServer
}

func (h *Handler) GetProduct(ctx context.Context, req *pbproduct.ProductRequest) (*pbproduct.ProductResponse, error) {
	return &pbproduct.ProductResponse{
		Id:    "1",
		Name:  "Product 1",
		Price: "1000",
	}, nil
}
