package handlers

import (
	"context"

	pb "github.com/sabnak227/jwt-demo/users"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.UserServer {
	return userService{}
}

type userService struct{}

// AuthUser implements Service.
func (s userService) AuthUser(ctx context.Context, in *pb.AuthUserRequest) (*pb.AuthUserResponse, error) {
	var resp pb.AuthUserResponse
	var code int32
	var msg string
	if in.Email == "jasonheshuai@gmail.com" && in.Password == "123qwe" {
		code = 1
		msg = "success"
	} else {
		code = 2
		msg = "fail"
	}
	resp = pb.AuthUserResponse{
		Code: code,
		Message: msg,
	}
	return &resp, nil
}
