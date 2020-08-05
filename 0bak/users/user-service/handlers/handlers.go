package handlers

import (
	"context"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	pb "github.com/sabnak227/jwt-demo/bak/users"
	"github.com/sabnak227/jwt-demo/bak/users/user-service/config"
	"github.com/sabnak227/jwt-demo/bak/users/user-service/models"
)

var dbconn models.DBClient

var logger *log.Entry

func init() {
	database := config.NewDatabaseConfigurations(
		config.GetConfigFromEnv("DB_DRIVER", "mysql"),
		config.GetConfigFromEnv("DB_HOST", "mysql"),
		config.GetConfigFromEnv("DB_PORT", "3306"),
		config.GetConfigFromEnv("DB_NAME", "users"),
		config.GetConfigFromEnv("DB_USER", "users"),
		config.GetConfigFromEnv("DB_PASSWORD", "users"),
	)

	c := config.NewConfigurations(
		config.GetConfigFromEnv("PORT", ":8080"),
		database,
	)

	logger = log.WithFields(log.Fields{"request_id": uuid.New().String()})

	{
		logger.Info("Database Initilizing...")
		dbconn = &models.MysqlClient{}
		err := dbconn.OpenCon(c.Database, logger)
		if err != nil {
			panic("Database initialization failed" + err.Error())
		}
		logger.Info("Database initialized")
		dbconn.Migrate()
		logger.Info("Database migrated")
	}

}

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.UserServer {
	return userService{
		db: dbconn,
	}
}

type userService struct {
	db models.DBClient
}

// ListUser implements Service.
func (s userService) ListUser(ctx context.Context, in *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	var resp pb.ListUserResponse
	resp = pb.ListUserResponse{
		Code:  1,
		Users: nil,
	}
	return &resp, nil
}

// GetUser implements Service.
func (s userService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var resp pb.GetUserResponse
	resp = pb.GetUserResponse{
		// Code:
		// User:
	}
	return &resp, nil
}

// CreateUser implements Service.
func (s userService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var resp pb.CreateUserResponse
	u := models.User{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Password:  in.Password,
		Address1:  in.Address1,
		Address2:  in.Address2,
		City:      in.City,
		State:     in.State,
		Country:   in.Country,
		Phone:     in.Phone,
	}
	log.Printf("Creating user %s", u.String())

	code := int32(0)
	if err := s.db.InsertUser(u); err != nil {
		code = 1
	}
	resp = pb.CreateUserResponse{
		Code: code,
	}
	return &resp, nil
}

// UpdateUser implements Service.
func (s userService) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	var resp pb.UpdateUserResponse
	resp = pb.UpdateUserResponse{
		// Code:
	}
	return &resp, nil
}

// DeleteUser implements Service.
func (s userService) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	var resp pb.DeleteUserResponse
	resp = pb.DeleteUserResponse{
		// Code:
	}
	return &resp, nil
}

// AuthUser implements Service.
func (s userService) AuthUser(ctx context.Context, in *pb.AuthUserRequest) (*pb.AuthUserResponse, error) {
	var resp pb.AuthUserResponse
	resp = pb.AuthUserResponse{
		// Code:
	}
	return &resp, nil
}
