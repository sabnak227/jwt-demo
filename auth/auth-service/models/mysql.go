package models

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sabnak227/jwt-demo/auth/auth-service/config"
	"golang.org/x/crypto/bcrypt"
)

// CassandraClient cassandra client
type MysqlClient struct {
	conn *gorm.DB
	config config.Config
	logger *log.Logger
}

// OpenCon opens a connection
func (c *MysqlClient) OpenCon(config config.Config, logger *log.Logger) error {
	conStr := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
	db, err := gorm.Open(config.DBDriver, conStr)
	if err != nil {
		return err
	}
	db.LogMode(true)
	db.SetLogger(NewGormLogger(logger))

	c.conn = db
	c.config = config
	c.logger = logger
	return nil
}

func (c *MysqlClient) Migrate() {
	if c.config.AutoMigrate {
		c.conn.AutoMigrate(&Auth{})
	}
}

func (c *MysqlClient) Close() error {
	err := c.conn.Close()
	return err
}

func (c *MysqlClient) AuthUser(email string, password string) *Auth {
	var user Auth
	c.conn.Where("email = ?", email).First(&user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) ; err == nil {
		return &user
	}
	return nil
}