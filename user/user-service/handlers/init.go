package handlers

import (
	"github.com/sabnak227/jwt-demo/auth"
	authClient "github.com/sabnak227/jwt-demo/auth/auth-service/svc/client/grpc"
	"github.com/sabnak227/jwt-demo/user/user-service/config"
	"github.com/sabnak227/jwt-demo/user/user-service/models"
	amqpAdapter "github.com/sabnak227/jwt-demo/util/amqp"
	"github.com/sabnak227/jwt-demo/util/helper"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	authSvc auth.AuthServer
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
		AuthSvcHost: helper.GetStrFromEnv("AUTH_SVC_HOST", "auth:5040"),
		AmqpDsn: helper.GetStrFromEnv("AMQP_DSN", "amqp://guest:guest@rabbitmq:5672/"),
	}

	setupDb()
	setupGrpcClient()
	setupAMQP()
}

func setupGrpcClient() {
	logger.Info("Dialing auth service rpc server...")
	uconn, err := grpc.Dial(conf.AuthSvcHost, grpc.WithInsecure())
	if err != nil {
		panic("failed to connect to auth svc " + err.Error())
	}
	authSvc, _ = authClient.New(uconn)
}

func setupDb() {
	logger.Info("Database Initializing...")
	repo = &models.MysqlClient{}
	err := repo.OpenCon(conf, logger)
	if err != nil {
		panic("Database initialization failed, " + err.Error())
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
	logger.Info("AMQP initialized")
}