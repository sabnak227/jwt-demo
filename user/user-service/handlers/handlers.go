package handlers

import (
	"context"
	"encoding/json"
	"github.com/sabnak227/jwt-demo/auth"
	"github.com/sabnak227/jwt-demo/user/user-service/models"
	amqpAdapter "github.com/sabnak227/jwt-demo/util/amqp"
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

	// create authentication entry
	//res, err := authSvc.CreateAuth(ctx, &auth.CreateAuthRequest{
	//	UserId:    uint64(user.ID),
	//	Password:  in.Password,
	//	Email:     user.Email,
	//	FirstName: user.FirstName,
	//	LastName:  user.LastName,
	//})
	//if res == nil || res.Code != constant.SuccessCode {
	//	// failed, rollback
	//	return &pb.CreateUserResponse{
	//		Code:    constant.FailCode,
	//		Message: "Failed to create authentication entry",
	//	}, nil
	//}


	// create authentication entry via amqp
	msg := models.CreateUserMsg{
		UserId:    uint64(user.ID),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  in.Password,
	}

	b, _ := json.Marshal(msg)

	// publishing on user_create.# topic exchange to notify all services subscribing to this topic
	o := amqpAdapter.TopicPublisher("user_create", "user_create.#")
	if err := amqpClient.Publish(*o, b, "user_create.#"); err != nil {
		logger.Error("Failed to publish to queue")
		if err := repo.Delete(uint64(user.ID)); err != nil {
			logger.Error("Failed to delete faulty user: %d", user.ID)
		}
	}

	return &pb.CreateUserResponse{
		Code:    constant.SuccessCode,
		Message: "success",
	}, nil
}

// DeleteUser implements Service.
func (s userService) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	logger.Infof("deleting user: %d", in.ID)

	if err := repo.Delete(in.ID); err != nil {
		return &pb.DeleteUserResponse{
			Code:    constant.FailCode,
			Message: "Failed to delete user",
		}, nil
	}

	// delete authentication entry
	// no need to check the response anyways
	res, _ := authSvc.DeleteAuth(ctx, &auth.DeleteAuthRequest{
		UserId:    in.ID,
	})
	if res.Code != constant.SuccessCode {
		logger.Errorf("deleting auth entity for user %d failed", in.ID)
	}

	return &pb.DeleteUserResponse{
		Code:    constant.SuccessCode,
		Message: "success",
	}, nil
}
