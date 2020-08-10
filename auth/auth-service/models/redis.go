package models

import (
	"context"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"

	"github.com/sabnak227/jwt-demo/auth/auth-service/config"
)

// RedisClient redis client
type RedisClient struct {
	conn *redis.Client
	config config.Config
	logger *log.Logger
}

var (
	ctx = context.Background()
)

// OpenCon opens a connection
func (c *RedisClient) OpenCon(config config.Config, logger *log.Logger) error {
	r := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})
	c.conn = r
	c.config = config
	c.logger = logger
	_, err := r.Ping(ctx).Result()
	return err
}

func (c *RedisClient) Close() error {
	err := c.conn.Close()
	return err
}
