package main

import (
	"crypto/rsa"
	"io/ioutil"
	"os"

	"github.com/go-kit/kit/log"

	"github.com/dgrijalva/jwt-go"
	"github.com/sabnak/jwt-demo/auth_svc/routers"
)

// location of the files used for signing and verification
const (
	privKeyPath = "keys/app.rsa"     // openssl genrsa -out app.rsa 2048
	pubKeyPath  = "keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

// keys are held in global variables
var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
	logger    log.Logger
)

// read the key files before starting http handlers
func init() {
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		logger.Log("error", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		logger.Log("error", err)
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		logger.Log("error", err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		logger.Log("error", err)
	}
}

// setup the handlers and start listening to requests
func main() {

	r := routers.NewRouter(signKey, logger)

	r.Run(":8081")
}
