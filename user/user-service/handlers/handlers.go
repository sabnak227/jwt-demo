package handlers

import (
	"context"
	"github.com/sabnak227/jwt-demo/user/user-service/models"
	"github.com/sabnak227/jwt-demo/util/constant"
	"github.com/sabnak227/jwt-demo/util/helper"

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
	u, err := repo.GetUser(in.ID)
	if err != nil {
		return &pb.GetUserResponse{
			Code:    constant.UserNotFound,
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

// CreateUser implements Service.
func (s userService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	i := createUserRequest{
		req: *in,
	}
	// request body validation
	if err := i.Validate(); err != nil {
		errors, _ := helper.BuildErrorResponse(err)
		return &pb.CreateUserResponse{
			Code:    constant.ValidationError,
			Message: "Validation error",
			Errors:  errors,
		}, nil
	}

	// verify if user already exists
	a, err := repo.CheckEmailExists(in.Email)
	if err != nil && a != nil {
		return &pb.CreateUserResponse{
			Code:    constant.ValidationError,
			Message: "user already exist",
		}, nil
	}
	logger.Infof("creating user: %s", in.Email)

	// create user
	user := models.User{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Address1:  in.Address1,
		Address2:  in.Address2,
		City:      in.City,
		State:     in.State,
		Country:   in.Country,
		Phone:     in.Phone,
	}

	if err := repo.CreateUser(user); err != nil {
		return &pb.CreateUserResponse{
			Code:    constant.FailCode,
			Message: "Failed to create user",
		}, nil
	}

	return &pb.CreateUserResponse{
		Code:    constant.SuccessCode,
		Message: "success",
	}, nil
}
