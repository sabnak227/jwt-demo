package handlers

import (
	"context"
	"github.com/sabnak227/jwt-demo/auth/auth-service/token"
	"github.com/sabnak227/jwt-demo/scope"
	user "github.com/sabnak227/jwt-demo/user"
	"google.golang.org/grpc"
	"log"
	"os"

	pb "github.com/sabnak227/jwt-demo/auth"
	scopeClient "github.com/sabnak227/jwt-demo/scope/scope-service/svc/client/grpc"
	userClient "github.com/sabnak227/jwt-demo/user/user-service/svc/client/grpc"
	"github.com/sabnak227/jwt-demo/util/constant"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.AuthServer {
	return authService{}
}

type authService struct{}

var (
	userSvc  user.UserServer
	scopeSvc scope.ScopeServer
)

// read the key files before starting http handlers
func init() {
	var userHost string
	var scopeHost string

	if addr := os.Getenv("USER_HOST"); addr != "" {
		userHost = addr
	}
	if addr := os.Getenv("SCOPE_HOST"); addr != "" {
		scopeHost = addr
	}

	uconn, err := grpc.Dial(userHost, grpc.WithInsecure())
	if err != nil {
		log.Printf("failed to connect to user svc %s", err.Error())
	}
	userSvc, _ = userClient.New(uconn)

	sconn, err := grpc.Dial(scopeHost, grpc.WithInsecure())
	if err != nil {
		log.Printf("failed to connect to scope svc %s", err.Error())
	}
	scopeSvc, _ = scopeClient.New(sconn)
}

// JWKS implements Service.
func (s authService) JWKS(ctx context.Context, in *pb.JWKSRequest) (*pb.JWKSResponse, error) {
	var resp pb.JWKSResponse
	jwk := token.GetJWk()
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

// AuthUser implements Service.

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
			Code: constant.WrongPasswordCode,
			Message: "Wrong email and password combination",
		}, nil
	}

	u, err := userSvc.GetUser(ctx, &user.GetUserRequest{
		ID:    uint64(a.ID),
	})

	if u == nil || u.Code != constant.SuccessCode {
		return &pb.LoginResponse{
			Code:    constant.FailCode,
			Message: "Failed retrieving user info",
		}, err
	}

	sc, err := scopeSvc.UserScope(ctx, &scope.UserScopeRequest{})
	if sc == nil {
		return &pb.LoginResponse{
			Code:    constant.FailCode,
			Message: "Failed retrieving user scope",
		}, err
	}
	scopes := sc.Scopes

	tokenDetail, err := token.GenToken(scopes, u, sc)

	if err != nil {
		log.Printf("Failed to sign token %s", err.Error())
		return &pb.LoginResponse{
			Code:    constant.FailCode,
			Message: "Failed generating auth token",
		}, err
	}

	return &pb.LoginResponse{
		Code:         constant.SuccessCode,
		Message:      "Success",
		AccessToken:  tokenDetail.AccessToken,
		RefreshToken: tokenDetail.RefreshToken,
	}, nil
}
