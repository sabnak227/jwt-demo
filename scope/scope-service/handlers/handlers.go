package handlers

import (
	"context"
	"github.com/sabnak227/jwt-demo/util/constant"

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
		Code: constant.SuccessCode,
		Message: "success",
		Scopes: []string{"whatever", "hola"},
	}
	return &resp, nil
}
