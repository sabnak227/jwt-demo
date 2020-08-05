package handlers

import (
	"context"
	"crypto/rsa"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	pb "github.com/sabnak227/jwt-demo/auth"
	uclient "github.com/sabnak227/jwt-demo/users/user-service/svc/client/grpc"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.AuthServer {
	return authService{}
}

type authService struct{}

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

const (
	privKeyPath = "auth-service/keys/app.rsa"     // openssl genrsa -out app.rsa 2048
	pubKeyPath  = "auth-service/keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

// read the key files before starting http handlers
func init() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Printf("error %s", err.Error())
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Printf("error %s", err.Error())
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Printf("error %s", err.Error())
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Printf("error %s", err.Error())
	}
	log.Printf("verifyKey %s", verifyKey)

	uconn, err := grpc.Dial(":5053")
	if err != nil {
		log.Printf("failed to connect to user svc %s", err.Error())
	}
	uclient.New(uconn)

	//sconn, err := grpc.Dial(":5063")
	//if err != nil {
	//	log.Printf("failed to connect to user svc %s", err.Error())
	//}
}

// Login implements Service.
func (s authService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	var resp pb.LoginResponse
	var code int32
	var msg string
	user := "jason"

	// create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("RS256"))
	claims := make(jwt.MapClaims)

	// set our claims
	claims["AccessToken"] = "level1"
	claims["CustomUserInfo"] = struct {
		Name string
		Kind string
	}{user, "human"}

	// set the expire time
	// see http://tools.ietf.org/html/draft-ietf-oauth-json-web-token-20#section-4.1.4
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	t.Claims = claims

	tokenString, err := t.SignedString(signKey)
	if err != nil {
		log.Printf("Failed to sign token %s", err.Error())
		msg = ""
		code = 1
	} else {
		msg = tokenString
		code = 2
	}

	resp = pb.LoginResponse{
		Code: code,
		Message: msg,
	}
	return &resp, nil
}
