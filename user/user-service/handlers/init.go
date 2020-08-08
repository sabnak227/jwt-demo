package handlers

import (
	"github.com/sabnak227/jwt-demo/user/user-service/config"
	"github.com/sabnak227/jwt-demo/user/user-service/models"
	log "github.com/sirupsen/logrus"
	"os"
)

var conf config.Config
var logger *log.Logger
var repo models.DBClient


func init() {
	logger = log.New()
	conf = config.Config{
		DBDriver: getConfigFromEnv("DB_DRIVER", "mysql").(string),
		DBHost: getConfigFromEnv("DB_HOST", "mysql").(string),
		DBPort: getConfigFromEnv("DB_PORT", "3306").(string),
		DBName: getConfigFromEnv("DB_NAME", "users").(string),
		DBUser: getConfigFromEnv("DB_USER", "users").(string),
		DBPassword: getConfigFromEnv("DB_PASSWORD", "users").(string),
		AutoMigrate: getConfigFromEnv("AUTO_MIGRATE", true).(bool),
	}

	{
		logger.Info("Database Initilizing...")
		repo = &models.MysqlClient{}
		err := repo.OpenCon(conf, logger)
		if err != nil {
			panic("Database initialization failed" + err.Error())
		}
		logger.Info("Database initialized")
		repo.Migrate()
		logger.Info("Database migrated")
	}
}

func getConfigFromEnv(key string, defaultVal interface{}) interface{} {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal
	} else {
		return val
	}
}