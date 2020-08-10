package token

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/sabnak227/jwt-demo/auth/auth-service/config"
	"github.com/sabnak227/jwt-demo/auth/auth-service/models"
	"github.com/sabnak227/jwt-demo/scope"
	"github.com/sabnak227/jwt-demo/user"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"time"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
	jwks		  Jwks
	accessExp = time.Minute * 15
	refreshExp = time.Hour * 24 * 7
)

const (
	pubKeyPathDefault  = "auth-service/keys/app.rsa.pub"
	privKeyPathDefault = "auth-service/keys/app.rsa"
	jwksPathDefault = "auth-service/keys/jwks.json"
)

func init() {
	pubKeyPath := pubKeyPathDefault
	privKeyPath := privKeyPathDefault
	jwksPath := jwksPathDefault

	if addr := os.Getenv("PUB_KEY_PATH"); addr != "" {
		pubKeyPath = addr
	}
	if addr := os.Getenv("PRI_KEY_PATH"); addr != "" {
		privKeyPath= addr
	}
	if addr := os.Getenv("JWKS_PATH"); addr != "" {
		jwksPath= addr
	}
	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		panic(fmt.Sprintf("error %s", err.Error()))
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic(fmt.Sprintf("error %s", err.Error()))
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		panic(fmt.Sprintf("error %s", err.Error()))
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		panic(fmt.Sprintf("error %s", err.Error()))
	}

	jwkBytes, err := ioutil.ReadFile(jwksPath)
	if err != nil {
		panic(fmt.Sprintf("error %s", err.Error()))
	}

	if err := json.Unmarshal(jwkBytes, &jwks); err != nil {
		panic(fmt.Sprintf("error %s", err.Error()))
	}
}

type Details struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type Jwks struct {
	Kty 	string `json:"kty"`
	N 		string `json:"n"`
	E 		string `json:"e"`
}

type Token struct {
	Conf config.Config
	Logger *log.Logger
	Session models.SessionClient
}

func NewToken(conf config.Config, logger *log.Logger, session models.SessionClient) *Token {
	return &Token{
		Conf: conf,
		Logger: logger,
		Session: session,
	}
}

func (t *Token)GenToken(scopes []string, user *user.GetUserResponse, scope *scope.UserScopeResponse) (*Details, error) {
	td := &Details{}
	td.AtExpires = time.Now().Add(accessExp).Unix()
	td.AccessUuid = uuid.New().String()

	td.RtExpires = time.Now().Add(refreshExp).Unix()
	td.RefreshUuid = uuid.New().String()
	var err error
	// create access token
	// TODO: use NewWithClaims method here
	at := jwt.New(jwt.GetSigningMethod("RS256"))
	atClaims := make(jwt.MapClaims)
	atClaims["iss"] = "admin@example.com"
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["scopes"] = scopes
	atClaims["exp"] = td.AtExpires
	atClaims["user_info"] = user
	atClaims["scope"] = scope
	at.Claims = atClaims
	td.AccessToken, err = at.SignedString(signKey)
	if err != nil {
		return nil, err
	}

	// create refresh token
	rt := jwt.New(jwt.GetSigningMethod("RS256"))
	rtClaims := make(jwt.MapClaims)
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["exp"] = td.RtExpires
	rt.Claims = rtClaims
	td.RefreshToken, err = rt.SignedString(signKey)
	if err != nil {
		return nil, err
	}

	return td, nil
}


func (t *Token)RefreshToken(refreshToken string) (*Details, error) {
	_, err := t.verifyToken(refreshToken)
	if err != nil {
		return nil, err
	}

	td := &Details{
	}
	return td, err
}

func (t *Token)verifyToken(refreshToken string) (interface{}, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (t *Token)GetJWk() Jwks{
	return jwks
}