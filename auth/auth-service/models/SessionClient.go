package models

import (
	"github.com/sabnak227/jwt-demo/auth/auth-service/config"
	"github.com/sabnak227/jwt-demo/auth/auth-service/token"
	"github.com/sabnak227/jwt-demo/user"
	log "github.com/sirupsen/logrus"
)

// SessionClient session connection interface
type SessionClient interface {
	OpenCon(config config.Config, logger *log.Logger) error
	Close() error
	SetToken(uint64, *token.Details, *user.UserObj, []string) error
	GetUserIdByRefreshUUID(string) (uint64, error)
	GetUserInfo(uint64) (*user.UserObj, []string, error)
	SetUserInfo(uint64, *user.UserObj, []string) error
}
