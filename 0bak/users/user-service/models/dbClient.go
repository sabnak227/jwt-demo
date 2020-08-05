package models

import (
	"github.com/sabnak227/jwt-demo/bak/users/user-service/config"
	log "github.com/sirupsen/logrus"
)

// DBClient db connection interface
type DBClient interface {
	OpenCon(config config.DatabaseConfigurations, logger *log.Entry) error
	Migrate()
	Close()
	GetUsers() []User
	InsertUser(User) error
}

// GormLogger struct
type GormLogger struct {
	Logger *log.Entry
}

func NewGormLogger(logger *log.Entry) *GormLogger {
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
