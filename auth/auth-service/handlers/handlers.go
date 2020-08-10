package handlers

import (
	"context"
	"fmt"
	pb "github.com/sabnak227/jwt-demo/auth"
	"github.com/sabnak227/jwt-demo/scope"
	"github.com/sabnak227/jwt-demo/user"
	"github.com/sabnak227/jwt-demo/util/constant"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.AuthServer {
	return authService{}
}

type authService struct{}

// JWKS implements Service.
func (s authService) JWKS(ctx context.Context, in *pb.JWKSRequest) (*pb.JWKSResponse, error) {
	var resp pb.JWKSResponse
	jwk := tokenAdapter.GetJWk()
	res := pb.JWKSResponse_Keys{
		Kty: jwk.Kty,
		N:   jwk.N,
		E:   jwk.E,
	}
	resp = pb.JWKSResponse{
		Keys: []*pb.JWKSResponse_Keys{&res},
	}
	return &resp, nil
}

// Login implements Service.
func (s authService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	logger.Infof("User %s is logging in", in.Email)
	i := loginRequest{
		req: *in,
	}

	if err := i.Validate(); err != nil {
		return nil, err
	}

	a := repo.AuthUser(in.Email, in.Password)
	if a == nil {
		return &pb.LoginResponse{
			Code:    constant.WrongPasswordCode,
			Message: "Wrong email and password combination",
		}, nil
	}

	u, err := userSvc.GetUser(ctx, &user.GetUserRequest{
		ID: uint64(a.ID),
	})

	if u == nil || u.Code != constant.SuccessCode {
		return &pb.LoginResponse{
			Code:    constant.UserNotFound,
			Message: "Failed retrieving user info",
		}, err
	}

	sc, err := scopeSvc.UserScope(ctx, &scope.UserScopeRequest{})
	if sc == nil {
		return &pb.LoginResponse{
			Code:    constant.ScopeNotFound,
			Message: "Failed retrieving user scope",
		}, err
	}
	scopes := sc.Scopes

	tokenDetail, err := tokenAdapter.GenToken(scopes, u, sc)

	if err != nil {
		return &pb.LoginResponse{
			Code:    constant.FailedGeneratingToken,
			Message: "Failed generating auth token",
		}, err
	}

	logger.Infof("Failed to sign token %v", tokenDetail)

	return &pb.LoginResponse{
		Code:         constant.SuccessCode,
		Message:      "Success",
		AccessToken:  tokenDetail.AccessToken,
		RefreshToken: tokenDetail.RefreshToken,
	}, nil
}

// Refresh implements Service.
func (s authService) Refresh(ctx context.Context, in *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	var resp pb.RefreshResponse
	tokenDetail, err := tokenAdapter.RefreshToken(in.RefreshToken)

	if err != nil {
		return &pb.RefreshResponse{
			Code: constant.FailCode,
			Message: fmt.Sprintf("Failed refresh token, %s", err),
		}, nil
	}

	resp = pb.RefreshResponse{
		Code: constant.SuccessCode,
		Message: "New access token granted",
		AccessToken:  tokenDetail.AccessToken,
		RefreshToken: tokenDetail.RefreshToken,
	}
	return &resp, nil
}
