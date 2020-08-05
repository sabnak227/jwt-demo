package main

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/sabnak/jwt-demo/user_svc/config"
	"github.com/sabnak/jwt-demo/user_svc/repository"
	"github.com/sabnak/jwt-demo/user_svc/routers"
)

var c *config.Configurations

func init() {
	database := config.NewDatabaseConfigurations(
		config.GetConfigFromEnv("DB_DRIVER", "mysql"),
		config.GetConfigFromEnv("DB_HOST", "mysql"),
		config.GetConfigFromEnv("DB_PORT", "3306"),
		config.GetConfigFromEnv("DB_NAME", "users"),
		config.GetConfigFromEnv("DB_USER", "users"),
		config.GetConfigFromEnv("DB_PASSWORD", "users"),
	)

	c = config.NewConfigurations(
		config.GetConfigFromEnv("PORT", ":8080"),
		database,
	)
}

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var dbconn repository.DBClient
	{
		logger.Log("msg", "Database Initilizing...")
		dbconn = &repository.MysqlClient{}
		err := dbconn.OpenCon(c.Database, logger)
		defer dbconn.Close()
		if err != nil {
			panic("Database initialization failed")
		}
		logger.Log("msg", "Database initialized")
		dbconn.Migrate()
		logger.Log("msg", "Database migrated")
	}

	r := routers.NewRouter(dbconn, logger)

	r.Run(c.Port)
}
