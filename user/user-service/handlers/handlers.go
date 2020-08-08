package handlers

import (
	"context"
	"github.com/sabnak227/jwt-demo/util/constant"

	pb "github.com/sabnak227/jwt-demo/user"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.UserServer {
	return userService{}
}

type userService struct{}

// GetUser implements Service.
func (s userService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	logger.Infof("Getting user info for %d", in.ID)
	var resp pb.GetUserResponse
	u := repo.GetUser(in.ID)
	if u == nil {
		return &pb.GetUserResponse{
			Code: constant.UserNotFound,
			Message: "User not found",
		}, nil
	}

	resp = pb.GetUserResponse{
		Code:    constant.SuccessCode,
		Message: "success",
		User: &pb.UserObj{
			Id:        uint64(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Address1:  u.Address1,
			Address2:  u.Address2,
			City:      u.City,
			State:     u.State,
			Country:   u.Country,
			Phone:     u.Phone,
		},
	}
	return &resp, nil
}
