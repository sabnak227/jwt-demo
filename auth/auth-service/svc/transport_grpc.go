// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 8907ffca23
// Version Date: Wed 27 Nov 2019 21:28:21 UTC

package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/sabnak227/jwt-demo/auth"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC AuthServer.
func MakeGRPCServer(endpoints Endpoints, options ...grpctransport.ServerOption) pb.AuthServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	serverOptions = append(serverOptions, options...)
	return &grpcServer{
		// auth

		jwks: grpctransport.NewServer(
			endpoints.JWKSEndpoint,
			DecodeGRPCJWKSRequest,
			EncodeGRPCJWKSResponse,
			serverOptions...,
		),
		login: grpctransport.NewServer(
			endpoints.LoginEndpoint,
			DecodeGRPCLoginRequest,
			EncodeGRPCLoginResponse,
			serverOptions...,
		),
		refresh: grpctransport.NewServer(
			endpoints.RefreshEndpoint,
			DecodeGRPCRefreshRequest,
			EncodeGRPCRefreshResponse,
			serverOptions...,
		),
		logout: grpctransport.NewServer(
			endpoints.LogoutEndpoint,
			DecodeGRPCLogoutRequest,
			EncodeGRPCLogoutResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the AuthServer interface
type grpcServer struct {
	jwks    grpctransport.Handler
	login   grpctransport.Handler
	refresh grpctransport.Handler
	logout  grpctransport.Handler
}

// Methods for grpcServer to implement AuthServer interface

func (s *grpcServer) JWKS(ctx context.Context, req *pb.JWKSRequest) (*pb.JWKSResponse, error) {
	_, rep, err := s.jwks.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.JWKSResponse), nil
}

func (s *grpcServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	_, rep, err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.LoginResponse), nil
}

func (s *grpcServer) Refresh(ctx context.Context, req *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	_, rep, err := s.refresh.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RefreshResponse), nil
}

func (s *grpcServer) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	_, rep, err := s.logout.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.LogoutResponse), nil
}

// Server Decode

// DecodeGRPCJWKSRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC jwks request to a user-domain jwks request. Primarily useful in a server.
func DecodeGRPCJWKSRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.JWKSRequest)
	return req, nil
}

// DecodeGRPCLoginRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC login request to a user-domain login request. Primarily useful in a server.
func DecodeGRPCLoginRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.LoginRequest)
	return req, nil
}

// DecodeGRPCRefreshRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC refresh request to a user-domain refresh request. Primarily useful in a server.
func DecodeGRPCRefreshRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RefreshRequest)
	return req, nil
}

// DecodeGRPCLogoutRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC logout request to a user-domain logout request. Primarily useful in a server.
func DecodeGRPCLogoutRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.LogoutRequest)
	return req, nil
}

// Server Encode

// EncodeGRPCJWKSResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain jwks response to a gRPC jwks reply. Primarily useful in a server.
func EncodeGRPCJWKSResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.JWKSResponse)
	return resp, nil
}

// EncodeGRPCLoginResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain login response to a gRPC login reply. Primarily useful in a server.
func EncodeGRPCLoginResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.LoginResponse)
	return resp, nil
}

// EncodeGRPCRefreshResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain refresh response to a gRPC refresh reply. Primarily useful in a server.
func EncodeGRPCRefreshResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.RefreshResponse)
	return resp, nil
}

// EncodeGRPCLogoutResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain logout response to a gRPC logout reply. Primarily useful in a server.
func EncodeGRPCLogoutResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.LogoutResponse)
	return resp, nil
}

// Helpers

func metadataToContext(ctx context.Context, md metadata.MD) context.Context {
	for k, v := range md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	return ctx
}
