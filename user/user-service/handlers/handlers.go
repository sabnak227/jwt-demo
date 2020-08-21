package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/sabnak227/jwt-demo/user"
	"github.com/sabnak227/jwt-demo/user/user-service/models"
	amqpAdapter "github.com/sabnak227/jwt-demo/util/amqp"
	"github.com/sabnak227/jwt-demo/util/constant"
	"github.com/sabnak227/jwt-demo/util/errors"
	"net/http"
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
	u, err := repo.GetUser(repo.GetConn(), in.ID)
	if err != nil {
		return nil, errors.NewResponseError(err, "User not found").SetErrorCode(constant.UserNotFound).SetStatusCode(http.StatusNotFound)
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
		return nil, errors.NewResponseError(err, "Validation error")
	}

	// verify if user already exists
	exist, err := repo.CheckEmailExists(repo.GetConn(), in.Email)
	if err != nil {
		return nil, errors.NewResponseError(err, "Failed getting user email")
	}
	if exist {
		return nil, errors.NewResponseError(fmt.Errorf("email exists"), "Email already taken").SetErrorCode(constant.ValidationError)
	}
	logger.Infof("creating user: %s", in.Email)

	// create user
	u := models.User{
		Status:    in.Status,
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

	user, err := repo.CreateUser(repo.GetConn(), u)
	if err != nil {
		return nil, errors.NewResponseError(err, "Failed to create user")
	}

	// publish user creation via amqp
	b, _ := json.Marshal(models.NewUserMsg(user, models.UserMsgTypeCreated).SetPassword(in.Password))

	// publishing on user_create.# fanout exchange to notify all services subscribing to this topic
	o := amqpAdapter.FanoutPublisher("user_updates")
	if err := amqpClient.Publish(*o, b, ""); err != nil {
		logger.Errorf("Failed to publish to queue %s", err)
	}

	return &pb.CreateUserResponse{
		Code:    constant.SuccessCode,
		Message: "success",
	}, nil
}

// UpdateUser implements Service.
func (s userService) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	i := updateUserRequest{
		req: *in,
	}
	// request body validation
	if err := i.Validate(); err != nil {
		return nil, errors.NewResponseError(err, "Validation error")
	}

	// get user
	user, err := repo.GetUser(repo.GetConn(), in.ID)
	if err != nil {
		return nil, errors.NewResponseError(err, "User not found").SetErrorCode(constant.UserNotFound).SetErrorCode(http.StatusNotFound)
	}

	// verify email
	if in.Email != user.Email {
		exist, err := repo.CheckEmailExists(repo.GetConn(), in.Email)
		if err != nil {
			return nil, errors.NewResponseError(err, "Failed getting user email")
		}
		if exist {
			return nil, errors.NewResponseError(fmt.Errorf("email exists"), "Email already taken").SetErrorCode(constant.ValidationError)
		}
	}

	logger.Infof("updating user: %s", in.Email)

	// update user
	user.FirstName = in.FirstName
	user.LastName = in.LastName
	user.Email = in.Email
	user.Address1 = in.Address1
	user.Address2 = in.Address2
	user.City = in.City
	user.State = in.State
	user.Country = in.Country
	user.Phone = in.Phone
	user.Status = in.Status

	if _, err := repo.UpdateUser(repo.GetConn(), *user); err != nil {
		return nil, errors.NewResponseError(err, "Failed to update user")
	}

	return &pb.UpdateUserResponse{
		Code:    constant.SuccessCode,
		Message: "success",
	}, nil
}

// DeleteUser implements Service.
func (s userService) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	logger.Infof("deleting user: %d", in.ID)

	user, err := repo.GetUser(repo.GetConn(), in.ID)
	if err != nil {
		return nil, errors.NewResponseError(err, "Failed to delete user")
	}

	if err := repo.Delete(repo.GetConn(), in.ID); err != nil {
		return nil, errors.NewResponseError(err, "Failed to create user")
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
