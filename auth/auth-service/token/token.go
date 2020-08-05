package token

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"io/ioutil"
	"time"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

const (
	privKeyPath = "auth-service/keys/app.rsa"
	pubKeyPath  = "auth-service/keys/app.rsa.pub"
)

func init() {
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
}

type Details struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func GenToken(scopes []string, userInfo interface{}) (*Details, error) {
	td := &Details{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.New().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.New().String()
	var err error
	// create access token
	at := jwt.New(jwt.GetSigningMethod("RS256"))
	atClaims := make(jwt.MapClaims)
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["scopes"] = scopes
	atClaims["exp"] = td.AtExpires
	atClaims["user_info"] = userInfo
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