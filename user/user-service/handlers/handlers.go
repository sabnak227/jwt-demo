package handlers

import (
	"context"
	"encoding/json"
	pb "github.com/sabnak227/jwt-demo/user"
	"github.com/sabnak227/jwt-demo/user/user-service/models"
	amqpAdapter "github.com/sabnak227/jwt-demo/util/amqp"
	"github.com/sabnak227/jwt-demo/util/constant"
	"github.com/sabnak227/jwt-demo/util/helper"
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
		logger.Error(err)
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
	exist, err := repo.CheckEmailExists(in.Email)
	if err == nil && exist != nil {
		return &pb.CreateUserResponse{
			Code:    constant.ValidationError,
			Message: "user already exist",
		}, nil
	}
	logger.Infof("creating user: %s", in.Email)

	// create user
	u := models.User{
		Status:    models.UserStatusEnabled,
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

	user, err := repo.CreateUser(u)
	if err != nil {
		return &pb.CreateUserResponse{
			Code:    constant.FailCode,
			Message: "Failed to create user",
		}, nil
	}

	// publish user creation via amqp
	b, _ := json.Marshal(models.NewUserMsg(user, models.UserMsgTypeCreated).SetPassword(in.Password))

	// publishing on user_create.# topic exchange to notify all services subscribing to this topic
	o := amqpAdapter.FanoutPublisher("user_updates")
	if err := amqpClient.Publish(*o, b, ""); err != nil {
		logger.Errorf("Failed to publish to queue %s", err)
	}

	return &pb.CreateUserResponse{
		Code:    constant.SuccessCode,
		Message: "success",
	}, nil
}

// DeleteUser implements Service.
func (s userService) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	logger.Infof("deleting user: %d", in.ID)

	user, err := repo.GetUser(in.ID)
	if err != nil {
		return &pb.DeleteUserResponse{
			Code:    constant.FailCode,
			Message: "Failed to delete user",
		}, nil
	}

	if err := repo.Delete(in.ID); err != nil {
		return &pb.DeleteUserResponse{
			Code:    constant.FailCode,
			Message: "Failed to delete user",
		}, nil
	}

	// publish user deletion via amqp
	b, _ := json.Marshal(models.NewUserMsg(user, models.UserMsgTypeDeleted))

	// publishing on user_create.# topic exchange to notify all services subscribing to this topic
	o := amqpAdapter.FanoutPublisher("user_updates")
	if err := amqpClient.Publish(*o, b, ""); err != nil {
		logger.Errorf("Failed to publish to queue %s", err)
	}

	return &pb.DeleteUserResponse{
		Code:    constant.SuccessCode,
		Message: "success",
	}, nil
}
