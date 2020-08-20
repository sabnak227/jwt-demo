package handlers

import (
	"github.com/sabnak227/jwt-demo/scope/scope-service/config"
	"github.com/sabnak227/jwt-demo/scope/scope-service/models"
	amqpAdapter "github.com/sabnak227/jwt-demo/util/amqp"
	"github.com/sabnak227/jwt-demo/util/helper"
	log "github.com/sirupsen/logrus"
)

var (
	conf config.Config
	logger *log.Logger
	repo models.DBClient
	amqpClient *amqpAdapter.AmqpClient
)

func init() {
	logger = log.New()

	conf = config.Config{
		DBDriver: helper.GetStrFromEnv("DB_DRIVER", "mysql"),
		DBHost: helper.GetStrFromEnv("DB_HOST", "mysql"),
		DBPort: helper.GetStrFromEnv("DB_PORT", "3306"),
		DBName: helper.GetStrFromEnv("DB_NAME", "users"),
		DBUser: helper.GetStrFromEnv("DB_USER", "users"),
		DBPassword: helper.GetStrFromEnv("DB_PASSWORD", "users"),
		AutoMigrate: helper.GetBoolFromEnv("AUTO_MIGRATE", true),
		Seed: helper.GetBoolFromEnv("SEED", false),
		AmqpDsn: helper.GetStrFromEnv("AMQP_DSN", "amqp://guest:guest@rabbitmq:5672/"),
	}
	setupDb()
	setupAMQP()
}

func setupDb() {
	logger.Info("Database Initializing...")
	repo = &models.MysqlClient{}
	err := repo.OpenCon(conf, logger)
	if err != nil {
		panic("Database initialization failed" + err.Error())
	}
	logger.Info("Database initialized")
	repo.Migrate()
	logger.Info("Database migrated")
}

func setupAMQP() {
	logger.Info("AMQP Initializing...")
	amqpClient = &amqpAdapter.AmqpClient{}
	err := amqpClient.ConnectToBroker(conf.AmqpDsn)
	if err != nil {
		panic("AMQP initialization failed" + err.Error())
	}
	subscribers()
	logger.Info("AMQP initialized")
}