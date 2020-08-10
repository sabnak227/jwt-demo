package models

import (
	"github.com/sabnak227/jwt-demo/auth/auth-service/config"
	log "github.com/sirupsen/logrus"
)

// SessionClient session connection interface
type SessionClient interface {
	OpenCon(config config.Config, logger *log.Logger) error
	Close() error
}
