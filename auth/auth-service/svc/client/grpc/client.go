// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 8907ffca23
// Version Date: Wed 27 Nov 2019 21:28:21 UTC

// Package grpc provides a gRPC client for the Auth service.
package grpc

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/sabnak227/jwt-demo/auth"
	"github.com/sabnak227/jwt-demo/auth/auth-service/svc"
)

// New returns an service backed by a gRPC client connection. It is the
// responsibility of the caller to dial, and later close, the connection.
func New(conn *grpc.ClientConn, options ...ClientOption) (pb.AuthServer, error) {
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
	var jwksEndpoint endpoint.Endpoint
	{
		jwksEndpoint = grpctransport.NewClient(
			conn,
			"auth.Auth",
			"JWKS",
			EncodeGRPCJWKSRequest,
			DecodeGRPCJWKSResponse,
			pb.JWKSResponse{},
			clientOptions...,
		).Endpoint()
	}

	var loginEndpoint endpoint.Endpoint
	{
		loginEndpoint = grpctransport.NewClient(
			conn,
			"auth.Auth",
			"Login",
			EncodeGRPCLoginRequest,
			DecodeGRPCLoginResponse,
			pb.LoginResponse{},
			clientOptions...,
		).Endpoint()
	}

	var createauthEndpoint endpoint.Endpoint
	{
		createauthEndpoint = grpctransport.NewClient(
			conn,
			"auth.Auth",
			"CreateAuth",
			EncodeGRPCCreateAuthRequest,
			DecodeGRPCCreateAuthResponse,
			pb.CreateAuthResponse{},
			clientOptions...,
		).Endpoint()
	}

	var refreshEndpoint endpoint.Endpoint
	{
		refreshEndpoint = grpctransport.NewClient(
			conn,
			"auth.Auth",
			"Refresh",
			EncodeGRPCRefreshRequest,
			DecodeGRPCRefreshResponse,
			pb.RefreshResponse{},
			clientOptions...,
		).Endpoint()
	}

	return svc.Endpoints{
		JWKSEndpoint:       jwksEndpoint,
		LoginEndpoint:      loginEndpoint,
		CreateAuthEndpoint: createauthEndpoint,
		RefreshEndpoint:    refreshEndpoint,
	}, nil
}

// GRPC Client Decode

// DecodeGRPCJWKSResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC jwks reply to a user-domain jwks response. Primarily useful in a client.
func DecodeGRPCJWKSResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.JWKSResponse)
	return reply, nil
}

// DecodeGRPCLoginResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC login reply to a user-domain login response. Primarily useful in a client.
func DecodeGRPCLoginResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.LoginResponse)
	return reply, nil
}

// DecodeGRPCCreateAuthResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC createauth reply to a user-domain createauth response. Primarily useful in a client.
func DecodeGRPCCreateAuthResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.CreateAuthResponse)
	return reply, nil
}

// DecodeGRPCRefreshResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC refresh reply to a user-domain refresh response. Primarily useful in a client.
func DecodeGRPCRefreshResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.RefreshResponse)
	return reply, nil
}

// GRPC Client Encode

// EncodeGRPCJWKSRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain jwks request to a gRPC jwks request. Primarily useful in a client.
func EncodeGRPCJWKSRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.JWKSRequest)
	return req, nil
}

// EncodeGRPCLoginRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain login request to a gRPC login request. Primarily useful in a client.
func EncodeGRPCLoginRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.LoginRequest)
	return req, nil
}

// EncodeGRPCCreateAuthRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain createauth request to a gRPC createauth request. Primarily useful in a client.
func EncodeGRPCCreateAuthRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateAuthRequest)
	return req, nil
}

// EncodeGRPCRefreshRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain refresh request to a gRPC refresh request. Primarily useful in a client.
func EncodeGRPCRefreshRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.RefreshRequest)
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
