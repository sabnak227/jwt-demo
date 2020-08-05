package routers

import (
	"crypto/rsa"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
	"github.com/sabnak/jwt-demo/auth_svc/controllers"
)

func NewRouter(signKey *rsa.PrivateKey, logger log.Logger) *gin.Engine {
	r := gin.New()
	r.Use(TimerMiddleware(logger))
	c := controllers.NewController(signKey, logger)

	r.GET("/login", c.Login)

	return r
}

func TimerMiddleware(logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		elapsed := end.Sub(start)
		logger.Log("elapsed time", elapsed.String())
	}
}
