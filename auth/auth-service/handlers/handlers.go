package handlers

import (
	"context"
	"fmt"
	pb "github.com/sabnak227/jwt-demo/auth"
	"github.com/sabnak227/jwt-demo/auth/auth-service/models"
	"github.com/sabnak227/jwt-demo/scope"
	"github.com/sabnak227/jwt-demo/user"
	"github.com/sabnak227/jwt-demo/util/constant"
	"github.com/sabnak227/jwt-demo/util/errors"
	"golang.org/x/crypto/bcrypt"
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
		return nil, errors.NewResponseError(err, "Validation error")
	}

	// verify user credentials in database
	a, err := repo.AuthUser(repo.GetConn(), in.Email, in.Password)
	if err != nil {
		return nil, errors.NewResponseError(err, "Wrong email and password combination").SetErrorCode(constant.WrongPasswordCode)
	}

	// get user info from other services
	u, sc, err := getUserInfo(ctx, a.UserID)
	if err != nil {
		return nil, errors.NewResponseError(err, "Failed getting user info")
	}

	// generate jwt token
	tokenDetail, err := tokenAdapter.GenToken(sc.Scopes, u.User, sc.Scopes)
	if err != nil {
		return nil, errors.NewResponseError(err, "Failed generating auth token")
	}

	// set session info
	if err := session.SetToken(a.UserID, tokenDetail, u.User, sc.Scopes); err != nil {
		return nil, errors.NewResponseError(err, "Failed to create session")
	}

	return &pb.LoginResponse{
		Code:         constant.SuccessCode,
		Message:      "Success",
		AccessToken:  tokenDetail.AccessToken,
		RefreshToken: tokenDetail.RefreshToken,
	}, nil
}

// CreateAuth implements Service.
func (s authService) CreateAuth(ctx context.Context, in *pb.CreateAuthRequest) (*pb.CreateAuthResponse, error) {
	i := createAuthRequest{
		req: *in,
	}
	// request body validation
	if err := i.Validate(); err != nil {
		return nil, errors.NewResponseError(err, "Validation error")
	}

	// hashing password
	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.NewResponseError(err, "Failed to hash password")
	}

	// store in database
	if err := repo.CreateAuth(repo.GetConn(), models.Auth{
		UserID:    in.UserId,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Password:  string(hash),
	}); err != nil {
		return nil, errors.NewResponseError(err, "Failed to create the authentication entry").SetErrorCode(constant.WrongPasswordCode)
	}

	return &pb.CreateAuthResponse{
		Code:    constant.SuccessCode,
		Message: "success",
	}, nil
}

// Refresh implements Service.
func (s authService) Refresh(ctx context.Context, in *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	i := refreshRequest{
		req: *in,
	}
	// request body validation
	if err := i.Validate(); err != nil {
		return nil, errors.NewResponseError(err, "Validation error")
	}

	// verify if token is valid or not
	refreshUUID, err := tokenAdapter.VerifyToken(in.RefreshToken)
	if err != nil {
		return nil, errors.NewResponseError(err, "Failed refresh token").SetErrorCode(constant.FailRefreshToken)
	}

	// get userid from session using uuid returned previously
	userID, err := session.GetUserIdByRefreshUUID(refreshUUID)
	if err != nil {
		return nil, errors.NewResponseError(err, "Session expired, please login again").SetErrorCode(constant.SessionExpired)
	}

	// get userinfo detail from cache
	u, sc, infoErr := session.GetUserInfo(userID)
	if infoErr != nil {
		logger.Errorf("Failed to get user info, recreating user info..., error: %s", err)
		// get user info from other services
		uRes, scRes, err := getUserInfo(ctx, userID)
		if err != nil || uRes == nil {
			return nil, errors.NewResponseError(infoErr, "Failed get user info")
		}
		u = uRes.User
		sc = scRes.Scopes
	}

	// generate new jwt token
	tokenDetail, err := tokenAdapter.GenToken(sc, u, sc)
	if err != nil {
		return nil, errors.NewResponseError(err, "Failed generating auth token")
	}

	// set session info
	if err := session.SetToken(userID, tokenDetail, u, sc); err != nil {
		return nil, errors.NewResponseError(err, "Failed to create session")
	}

	return &pb.RefreshResponse{
		Code:         constant.SuccessCode,
		Message:      "New access token granted",
		AccessToken:  tokenDetail.AccessToken,
		RefreshToken: tokenDetail.RefreshToken,
	}, nil
}

func getUserInfo(ctx context.Context, userID uint64) (*user.GetUserResponse, *scope.UserScopeResponse, error) {
	u, err := userSvc.GetUser(ctx, &user.GetUserRequest{
		ID: userID,
	})
	logger.Infof("user svc response %v, error: %s", u, err)
	if u == nil || u.Code != constant.SuccessCode {
		return nil, nil, fmt.Errorf("failed retrieving user info")
	}

	sc, err := scopeSvc.UserScope(ctx, &scope.UserScopeRequest{})
	logger.Infof("scope svc response %v, error: %s", sc, err)
	if sc == nil || sc.Code != constant.SuccessCode {
		return nil, nil, fmt.Errorf("failed retrieving user scope")
	}
	return u, sc, nil
}
