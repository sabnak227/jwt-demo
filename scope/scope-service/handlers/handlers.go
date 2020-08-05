package handlers

import (
	"context"

	pb "github.com/sabnak227/jwt-demo/scope"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.ScopeServer {
	return scopeService{}
}

type scopeService struct{}

// UserScope implements Service.
func (s scopeService) UserScope(ctx context.Context, in *pb.UserScopeRequest) (*pb.UserScopeResponse, error) {
	var resp pb.UserScopeResponse
	resp = pb.UserScopeResponse{
		Code: 1,
		Message: "success",
		Scopes: []string{"whatever", "hola"},
	}
	return &resp, nil
}
