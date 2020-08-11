package models

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sabnak227/jwt-demo/auth/auth-service/config"
	"github.com/sabnak227/jwt-demo/auth/auth-service/token"
	"github.com/sabnak227/jwt-demo/scope"
	"github.com/sabnak227/jwt-demo/user"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

// RedisClient redis client
type RedisClient struct {
	conn *redis.Client
	config config.Config
	logger *log.Logger
}

var (
	ctx = context.Background()
	userInfoKey = "userinfo:"
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

func (c *RedisClient) SetToken(td *token.Details, user *user.GetUserResponse, scope *scope.UserScopeResponse) error {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := c.conn.Set(ctx, td.AccessUuid, user.User.Id, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := c.conn.Set(ctx, td.RefreshUuid, user.User.Id, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}

	if err := c.SetUserInfo(user.User.Id, user, scope); err != nil {
		return err
	}

	return nil
}

func (c *RedisClient) GetUserIdByRefreshUUID(uuid string) (uint64, error) {
	str, err := c.conn.Get(ctx, uuid).Result()
	if err != nil {
		return 0, err
	}
	n, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%d of type %T", n, n)
	}
	return n, nil
}

func (c *RedisClient) GetUserInfo(userID uint64) (*user.GetUserResponse, *scope.UserScopeResponse, error) {
	userStr, userErr := c.conn.Get(ctx, getUserInfoUserKey(userID)).Result()
	if userErr != nil {
		return nil, nil, userErr
	}

	scopeStr, scopeErr := c.conn.Get(ctx, getUserInfoScopeKey(userID)).Result()
	if scopeErr != nil {
		return nil, nil, scopeErr
	}

	var userRes *user.GetUserResponse
	if err := json.Unmarshal([]byte(userStr), &userRes); err != nil {
		return nil, nil, err
	}

	var scopeRes *scope.UserScopeResponse
	if err := json.Unmarshal([]byte(scopeStr), &scopeRes); err != nil {
		return nil, nil, err
	}

	return userRes, scopeRes, nil
}

func (c *RedisClient) SetUserInfo(userID uint64, user *user.GetUserResponse, scope *scope.UserScopeResponse) error {
	u, err := json.Marshal(user)
	if err != nil {
		return err
	}

	s, err := json.Marshal(scope)
	if err != nil {
		return err
	}

	if err := c.conn.Set(ctx, getUserInfoUserKey(userID), u, time.Hour * 24).Err(); err != nil {
		return err
	}

	if err := c.conn.Set(ctx, getUserInfoScopeKey(userID), s, time.Hour * 24).Err(); err != nil {
		return err
	}
	return nil
}

func getUserInfoUserKey (userID uint64) string {
	return fmt.Sprintf("%s:%d:user", userInfoKey, userID)
}
func getUserInfoScopeKey (userID uint64) string {
	return fmt.Sprintf("%s:%d:scope", userInfoKey, userID)
}