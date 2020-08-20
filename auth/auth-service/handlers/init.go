package handlers

import (
	"github.com/sabnak227/jwt-demo/auth/auth-service/config"
	"github.com/sabnak227/jwt-demo/auth/auth-service/models"
	"github.com/sabnak227/jwt-demo/auth/auth-service/token"
	"github.com/sabnak227/jwt-demo/scope"
	scopeClient "github.com/sabnak227/jwt-demo/scope/scope-service/svc/client/grpc"
	"github.com/sabnak227/jwt-demo/user"
	userClient "github.com/sabnak227/jwt-demo/user/user-service/svc/client/grpc"
	amqpAdapter "github.com/sabnak227/jwt-demo/util/amqp"
	"github.com/sabnak227/jwt-demo/util/helper"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	userSvc  user.UserServer
	scopeSvc scope.ScopeServer
	conf config.Config
	logger *log.Logger
	repo models.DBClient
	session models.SessionClient
	tokenAdapter *token.Token
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
		RedisHost: helper.GetStrFromEnv("REDIS_HOST", "redis:6379"),
		RedisPassword: helper.GetStrFromEnv("REDIS_PASSWORD", ""),
		RedisDB: helper.GetIntFromEnv("REDIS_DB", 0),
		UserSvcHost: helper.GetStrFromEnv("USER_SVC_HOST", "user:5040"),
		ScopeSvcHost: helper.GetStrFromEnv("SCOPE_SVC_HOST", "scope:5040"),
		AmqpDsn: helper.GetStrFromEnv("AMQP_DSN", "amqp://guest:guest@rabbitmq:5672/"),
	}
	setupDb()
	setupRedis()
	setupGrpcClient()
	setUpTokenAdapter()
	setupAMQP()
}

func setupGrpcClient() {
	logger.Info("Dialing user service rpc server...")
	uconn, err := grpc.Dial(conf.UserSvcHost, grpc.WithInsecure())
	if err != nil {
		panic("failed to connect to user svc " + err.Error())
	}
	userSvc, _ = userClient.New(uconn)

	logger.Info("Dialing scope service rpc server...")
	sconn, err := grpc.Dial(conf.ScopeSvcHost, grpc.WithInsecure())
	if err != nil {
		panic("failed to connect to scope svc " + err.Error())
	}
	scopeSvc, _ = scopeClient.New(sconn)
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

func setupRedis() {
	logger.Info("Redis Initializing...")

	session = &models.RedisClient{}
	err := session.OpenCon(conf, logger)
	if err != nil {
		panic("Redis initialization failed" + err.Error())
	}
	logger.Info("Redis initialized")
}

func setUpTokenAdapter() {
	tokenAdapter = token.NewToken(conf, logger)
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