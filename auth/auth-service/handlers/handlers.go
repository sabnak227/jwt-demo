package handlers

import (
	"context"
	"github.com/sabnak227/jwt-demo/auth/auth-service/token"
	"github.com/sabnak227/jwt-demo/scope"
	user "github.com/sabnak227/jwt-demo/users"
	"google.golang.org/grpc"
	"log"
	"os"

	pb "github.com/sabnak227/jwt-demo/auth"
	scopeClient "github.com/sabnak227/jwt-demo/scope/scope-service/svc/client/grpc"
	userClient "github.com/sabnak227/jwt-demo/users/user-service/svc/client/grpc"
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

// Login implements Service.
func (s authService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	res, err := userSvc.AuthUser(ctx, &user.AuthUserRequest{
		Email:    in.Email,
		Password: in.Password,
	})
	if res == nil || res.Code != 1 {
		return &pb.LoginResponse{
			Code:    2,
			Message: "failed wrong password",
		}, err
	}

	res1, err := scopeSvc.UserScope(ctx, &scope.UserScopeRequest{})
	if res1 == nil {
		return &pb.LoginResponse{
			Code:    2,
			Message: "failed",
		}, err
	}
	scopes := res1.Scopes

	u := struct {
		Name string
	}{in.Email}

	tokenDetail, err := token.GenToken(scopes, u)

	if err != nil {
		log.Printf("Failed to sign token %s", err.Error())
		return &pb.LoginResponse{
			Code:    2,
			Message: "failed",
		}, err
	}

	return &pb.LoginResponse{
		Code:         1,
		Message:      "success",
		AccessToken:  tokenDetail.AccessToken,
		RefreshToken: tokenDetail.RefreshToken,
	}, nil
}

// JWKS implements Service.
func (s authService) JWKS(ctx context.Context, in *pb.JWKSRequest) (*pb.JWKSResponse, error) {
	var resp pb.JWKSResponse
	jwk := token.GetJWk()
	res := pb.JWKSResponse_Keys{
		Kty: jwk.Kty,
		N: jwk.N,
		E: jwk.E,
	}
	resp = pb.JWKSResponse{
		Keys: []*pb.JWKSResponse_Keys{&res},
	}
	return &resp, nil
}
