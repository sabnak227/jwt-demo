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
	pb "github.com/sabnak227/jwt-demo/user"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC UserServer.
func MakeGRPCServer(endpoints Endpoints, options ...grpctransport.ServerOption) pb.UserServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	serverOptions = append(serverOptions, options...)
	return &grpcServer{
		// user

		getuser: grpctransport.NewServer(
			endpoints.GetUserEndpoint,
			DecodeGRPCGetUserRequest,
			EncodeGRPCGetUserResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the UserServer interface
type grpcServer struct {
	getuser grpctransport.Handler
}

// Methods for grpcServer to implement UserServer interface

func (s *grpcServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	_, rep, err := s.getuser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUserResponse), nil
}

// Server Decode

// DecodeGRPCGetUserRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC getuser request to a user-domain getuser request. Primarily useful in a server.
func DecodeGRPCGetUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetUserRequest)
	return req, nil
}

// Server Encode

// EncodeGRPCGetUserResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain getuser response to a gRPC getuser reply. Primarily useful in a server.
func EncodeGRPCGetUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GetUserResponse)
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
