package controllers

import (
	"crypto/rsa"

	"github.com/go-kit/kit/log"
)

// CassandraClient cassandra client
type Controller struct {
	signKey *rsa.PrivateKey
	logger  log.Logger
}

func NewController(signKey *rsa.PrivateKey, logger log.Logger) *Controller {
	return &Controller{
		signKey: signKey,
		logger:  logger,
	}
}
