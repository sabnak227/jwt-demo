package handlers

import (
	"context"
	"github.com/sabnak227/jwt-demo/util/constant"
	"github.com/sabnak227/jwt-demo/util/errors"

	pb "github.com/sabnak227/jwt-demo/scope"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.ScopeServer {
	return scopeService{}
}

type scopeService struct{}

// UserScope implements Service.
func (s scopeService) UserScope(ctx context.Context, in *pb.UserScopeRequest) (*pb.UserScopeResponse, error) {
	perms, err := repo.GetPerms(repo.GetConn(), in.ID)
	if err != nil {
		return nil, errors.NewResponseError(err, "Failed getting user scopes")
	}

	return &pb.UserScopeResponse{
		Code:    constant.SuccessCode,
		Message: "success",
		Scopes:  perms,
	}, nil
}
