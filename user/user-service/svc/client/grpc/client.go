// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 8907ffca23
// Version Date: Wed 27 Nov 2019 21:28:21 UTC

// Package grpc provides a gRPC client for the User service.
package grpc

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/sabnak227/jwt-demo/user"
	"github.com/sabnak227/jwt-demo/user/user-service/svc"
)

// New returns an service backed by a gRPC client connection. It is the
// responsibility of the caller to dial, and later close, the connection.
func New(conn *grpc.ClientConn, options ...ClientOption) (pb.UserServer, error) {
	var cc clientConfig

	for _, f := range options {
		err := f(&cc)
		if err != nil {
			return nil, errors.Wrap(err, "cannot apply option")
		}
	}

	clientOptions := []grpctransport.ClientOption{
		grpctransport.ClientBefore(
			contextValuesToGRPCMetadata(cc.headers)),
	}
	var getuserEndpoint endpoint.Endpoint
	{
		getuserEndpoint = grpctransport.NewClient(
			conn,
			"user.User",
			"GetUser",
			EncodeGRPCGetUserRequest,
			DecodeGRPCGetUserResponse,
			pb.GetUserResponse{},
			clientOptions...,
		).Endpoint()
	}

	var createuserEndpoint endpoint.Endpoint
	{
		createuserEndpoint = grpctransport.NewClient(
			conn,
			"user.User",
			"CreateUser",
			EncodeGRPCCreateUserRequest,
			DecodeGRPCCreateUserResponse,
			pb.CreateUserResponse{},
			clientOptions...,
		).Endpoint()
	}

	var updateuserEndpoint endpoint.Endpoint
	{
		updateuserEndpoint = grpctransport.NewClient(
			conn,
			"user.User",
			"UpdateUser",
			EncodeGRPCUpdateUserRequest,
			DecodeGRPCUpdateUserResponse,
			pb.UpdateUserResponse{},
			clientOptions...,
		).Endpoint()
	}

	var deleteuserEndpoint endpoint.Endpoint
	{
		deleteuserEndpoint = grpctransport.NewClient(
			conn,
			"user.User",
			"DeleteUser",
			EncodeGRPCDeleteUserRequest,
			DecodeGRPCDeleteUserResponse,
			pb.DeleteUserResponse{},
			clientOptions...,
		).Endpoint()
	}

	return svc.Endpoints{
		GetUserEndpoint:    getuserEndpoint,
		CreateUserEndpoint: createuserEndpoint,
		UpdateUserEndpoint: updateuserEndpoint,
		DeleteUserEndpoint: deleteuserEndpoint,
	}, nil
}

// GRPC Client Decode

// DecodeGRPCGetUserResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC getuser reply to a user-domain getuser response. Primarily useful in a client.
func DecodeGRPCGetUserResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetUserResponse)
	return reply, nil
}

// DecodeGRPCCreateUserResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC createuser reply to a user-domain createuser response. Primarily useful in a client.
func DecodeGRPCCreateUserResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.CreateUserResponse)
	return reply, nil
}

// DecodeGRPCUpdateUserResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC updateuser reply to a user-domain updateuser response. Primarily useful in a client.
func DecodeGRPCUpdateUserResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.UpdateUserResponse)
	return reply, nil
}

// DecodeGRPCDeleteUserResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC deleteuser reply to a user-domain deleteuser response. Primarily useful in a client.
func DecodeGRPCDeleteUserResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.DeleteUserResponse)
	return reply, nil
}

// GRPC Client Encode

// EncodeGRPCGetUserRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain getuser request to a gRPC getuser request. Primarily useful in a client.
func EncodeGRPCGetUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetUserRequest)
	return req, nil
}

// EncodeGRPCCreateUserRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain createuser request to a gRPC createuser request. Primarily useful in a client.
func EncodeGRPCCreateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateUserRequest)
	return req, nil
}

// EncodeGRPCUpdateUserRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain updateuser request to a gRPC updateuser request. Primarily useful in a client.
func EncodeGRPCUpdateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateUserRequest)
	return req, nil
}

// EncodeGRPCDeleteUserRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain deleteuser request to a gRPC deleteuser request. Primarily useful in a client.
func EncodeGRPCDeleteUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteUserRequest)
	return req, nil
}

type clientConfig struct {
	headers []string
}

// ClientOption is a function that modifies the client config
type ClientOption func(*clientConfig) error

func CtxValuesToSend(keys ...string) ClientOption {
	return func(o *clientConfig) error {
		o.headers = keys
		return nil
	}
}

func contextValuesToGRPCMetadata(keys []string) grpctransport.ClientRequestFunc {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		var pairs []string
		for _, k := range keys {
			if v, ok := ctx.Value(k).(string); ok {
				pairs = append(pairs, k, v)
			}
		}

		if pairs != nil {
			*md = metadata.Join(*md, metadata.Pairs(pairs...))
		}

		return ctx
	}
}
