package controllers

import (
	"github.com/go-kit/kit/log"

	"github.com/sabnak/jwt-demo/user_svc/repository"
)

// CassandraClient cassandra client
type Controller struct {
	dbClient repository.DBClient
	logger   log.Logger
}

func NewController(dbClient repository.DBClient, logger log.Logger) *Controller {
	return &Controller{
		dbClient: dbClient,
		logger:   logger,
	}
}
