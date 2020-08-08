package handlers

import (
	"context"
	"github.com/sabnak227/jwt-demo/util/constant"

	pb "github.com/sabnak227/jwt-demo/users"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.UserServer {
	return userService{}
}

type userService struct{}

// GetUser implements Service.
func (s userService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var resp pb.GetUserResponse
	resp = pb.GetUserResponse{
		Code: constant.SuccessCode,
		Message: "success",
		User: &pb.UserObj{
			Id:        0,
			FirstName: "kjn",
			LastName:  "lkn",
			Email:     "lknj",
			Address1:  "lk",
			Address2:  "",
			City:      "",
			State:     "",
			Country:   "",
			Phone:     "",
		},
	}
	return &resp, nil
}
