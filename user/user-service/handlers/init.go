package handlers

import (
	"github.com/sabnak227/jwt-demo/auth"
	"github.com/sabnak227/jwt-demo/user/user-service/config"
	"github.com/sabnak227/jwt-demo/user/user-service/models"
	authClient "github.com/sabnak227/jwt-demo/auth/auth-service/svc/client/grpc"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"os"
)

var (
	authSvc auth.AuthServer
	conf config.Config
	logger *log.Logger
	repo models.DBClient
)



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
		AuthSvcHost: getConfigFromEnv("AUTH_SVC_HOST", "auth:5040").(string),
	}

	setupDb()
	setupGrpcClient()
}

func setupGrpcClient() {
	logger.Info("Dialing user service rpc server...")
	uconn, err := grpc.Dial(conf.AuthSvcHost, grpc.WithInsecure())
	if err != nil {
		panic("failed to connect to auth svc " + err.Error())
	}
	auth, _ = authClient.New(uconn)
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

func getConfigFromEnv(key string, defaultVal interface{}) interface{} {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal
	} else {
		return val
	}
}