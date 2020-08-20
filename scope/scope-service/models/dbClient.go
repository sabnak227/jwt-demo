package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sabnak227/jwt-demo/scope/scope-service/config"
	log "github.com/sirupsen/logrus"
)

// DBClient db connection interface
type DBClient interface {
	OpenCon(config config.Config, logger *log.Logger) error
	GetConn() *gorm.DB
	Migrate()
	Close() error
	Seed()
	GetPerms(conn *gorm.DB, userID uint64) ([]string, error)
	CreateRole(conn *gorm.DB, role Role) (*Role, error)
	CreatePermission(conn *gorm.DB, perm Permission) (*Permission, error)
	AttachPerm(conn *gorm.DB, role Role, perm Permission) error
	DetachPerm(conn *gorm.DB, role Role, perm Permission) error
	AssignRole(conn *gorm.DB, userID uint64, roleName string) error
}

// GormLogger struct
type GormLogger struct {
	Logger *log.Logger
}

func NewGormLogger(logger *log.Logger) *GormLogger {
	return &GormLogger{
		Logger: logger,
	}
}

// Print - Log Formatter
func (*GormLogger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":
		log.WithFields(
			log.Fields{
				"module":        "gorm",
				"type":          "sql",
				"rows_returned": v[5],
				"src":           v[1],
				"values":        v[4],
				"duration":      v[2],
			},
		).Info(v[3])
	case "log":
		log.WithFields(log.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}
