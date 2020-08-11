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

	// request body validation
	if err := i.Validate(); err != nil {
		return nil, err
	}

	// verify user credentials in database
	a := repo.AuthUser(in.Email, in.Password)
	if a == nil {
		return &pb.LoginResponse{
			Code:    constant.WrongPasswordCode,
			Message: "Wrong email and password combination",
		}, nil
	}

	// get user info from other services
	u, sc, err := getUserInfo(ctx, uint64(a.ID))
	if err != nil {
		return &pb.LoginResponse{
			Code:    constant.FailCode,
			Message: err.Error(),
		}, nil
	}

	// generate jwt token
	tokenDetail, err := tokenAdapter.GenToken(sc.Scopes, u, sc)
	if err != nil {
		return &pb.LoginResponse{
			Code:    constant.FailCode,
			Message: "Failed generating auth token",
		}, nil
	}

	// set session info
	if err := session.SetToken(tokenDetail, u, sc); err != nil {
		return &pb.LoginResponse{
			Code:    constant.FailCode,
			Message: fmt.Sprintf("Failed to create session, err: %s", err),
		}, nil
	}


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
	// verify if token is valid or not
	refreshUUID, err := tokenAdapter.VerifyToken(in.RefreshToken)
	if err != nil {
		return &pb.RefreshResponse{
			Code: constant.FailCode,
			Message: fmt.Sprintf("Failed refresh token, %s", err),
		}, nil
	}

	// get userid from session using uuid returned previously
	userID , err := session.GetUserIdByRefreshUUID(refreshUUID)
	if err != nil {
		logger.Infof("err %s", err)
		return &pb.RefreshResponse{
			Code: constant.FailCode,
			Message: "Session expiried, please login again",
		}, nil
	}

	// get userinfo detail from cache
	u, sc, infoErr := session.GetUserInfo(userID)
	if infoErr != nil {
		logger.Errorf("Failed to get user info, recreating user info..., error: %s", err)
		// get user info from other services
		u, sc, infoErr = getUserInfo(ctx, userID)
		if infoErr != nil {
			return &pb.RefreshResponse{
				Code: constant.FailCode,
				Message: infoErr.Error(),
			}, nil
		}
	}

	// generate new jwt token
	tokenDetail, err := tokenAdapter.GenToken(sc.Scopes, u, sc)
	if err != nil {
		return &pb.RefreshResponse{
			Code:    constant.FailCode,
			Message: "Failed generating auth token",
		}, nil
	}

	// set session info
	if err := session.SetToken(tokenDetail, u, sc); err != nil {
		return &pb.RefreshResponse{
			Code:    constant.FailCode,
			Message: "Failed to create session",
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


func getUserInfo(ctx context.Context, userID uint64) (*user.GetUserResponse, *scope.UserScopeResponse, error) {
	u, _ := userSvc.GetUser(ctx, &user.GetUserRequest{
		ID: userID,
	})

	if u == nil || u.Code != constant.SuccessCode {
		return nil, nil, fmt.Errorf("failed retrieving user info")
	}

	sc, _ := scopeSvc.UserScope(ctx, &scope.UserScopeRequest{})
	if sc == nil || sc.Code != constant.SuccessCode {
		return nil, nil, fmt.Errorf("failed retrieving user scope")
	}
	return u, sc, nil
}