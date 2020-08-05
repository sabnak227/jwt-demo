// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 8907ffca23
// Version Date: Wed 27 Nov 2019 21:28:21 UTC

// Package grpc provides a gRPC client for the Scope service.
package grpc

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/sabnak227/jwt-demo/scope"
	"github.com/sabnak227/jwt-demo/scope/scope-service/svc"
)

// New returns an service backed by a gRPC client connection. It is the
// responsibility of the caller to dial, and later close, the connection.
func New(conn *grpc.ClientConn, options ...ClientOption) (pb.ScopeServer, error) {
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
	var userscopeEndpoint endpoint.Endpoint
	{
		userscopeEndpoint = grpctransport.NewClient(
			conn,
			"scope.Scope",
			"UserScope",
			EncodeGRPCUserScopeRequest,
			DecodeGRPCUserScopeResponse,
			pb.UserScopeResponse{},
			clientOptions...,
		).Endpoint()
	}

	return svc.Endpoints{
		UserScopeEndpoint: userscopeEndpoint,
	}, nil
}

// GRPC Client Decode

// DecodeGRPCUserScopeResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC userscope reply to a user-domain userscope response. Primarily useful in a client.
func DecodeGRPCUserScopeResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.UserScopeResponse)
	return reply, nil
}

// GRPC Client Encode

// EncodeGRPCUserScopeRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain userscope request to a gRPC userscope request. Primarily useful in a client.
func EncodeGRPCUserScopeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UserScopeRequest)
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
