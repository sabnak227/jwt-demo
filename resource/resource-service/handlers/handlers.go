package handlers

import (
	"context"

	pb "github.com/sabnak227/jwt-demo/resource"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.ResourceServer {
	return resourceService{}
}

type resourceService struct{}

// GetResource implements Service.
func (s resourceService) GetResource(ctx context.Context, in *pb.ResourceRequest) (*pb.ResourceResponse, error) {
	var resp pb.ResourceResponse
	resp = pb.ResourceResponse{
		Code: 1,
		Message: "hola, Is this going to work?????????????",
	}
	return &resp, nil
}
