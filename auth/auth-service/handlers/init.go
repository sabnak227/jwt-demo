package handlers

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sabnak227/jwt-demo/auth/auth-service/config"
	"github.com/sabnak227/jwt-demo/auth/auth-service/models"
	"github.com/sabnak227/jwt-demo/scope"
	scopeClient "github.com/sabnak227/jwt-demo/scope/scope-service/svc/client/grpc"
	"github.com/sabnak227/jwt-demo/user"
	userClient "github.com/sabnak227/jwt-demo/user/user-service/svc/client/grpc"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"os"
)

var (
	userSvc  user.UserServer
	scopeSvc scope.ScopeServer
	conf config.Config
	logger *log.Logger
	repo models.DBClient
	redisClient *redis.Client
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
		RedisHost: getConfigFromEnv("REDIS_HOST", "localhost:6379").(string),
		RedisPassword: getConfigFromEnv("REDIS_PASSWORD", "").(string),
		RedisDB: getConfigFromEnv("REDIS_DB", 0).(int),
		AutoMigrate: getConfigFromEnv("AUTO_MIGRATE", true).(bool),
		UserSvcHost: getConfigFromEnv("USER_SVC_HOST", true).(string),
		ScopeSvcHost: getConfigFromEnv("SCOPE_SVC_HOST", true).(string),
	}
	setupDb()
	setupRedis()
	setupGrpcClient()
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

// read the key files before starting http handlers
func setupRedis() {
	logger.Info("Redis Initializing...")
	redisClient = redis.NewClient(&redis.Options{
		Addr:     conf.RedisHost,
		Password: conf.RedisPassword,
		DB:       conf.RedisDB,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic("Redis initialization failed" + err.Error())
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