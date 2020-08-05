package routers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
	"github.com/sabnak/jwt-demo/user_svc/controllers"
	"github.com/sabnak/jwt-demo/user_svc/repository"
)

func NewRouter(dbconn repository.DBClient, logger log.Logger) *gin.Engine {
	r := gin.New()
	r.Use(TimerMiddleware(logger))
	c := controllers.NewController(dbconn, logger)

	r.GET("/users", c.GetUsers)
	r.POST("/users", c.InsertUser)

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
