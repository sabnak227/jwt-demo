package models

import (
	"fmt"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sabnak227/jwt-demo/auth/auth-service/config"
	"golang.org/x/crypto/bcrypt"
)

// MysqlClient mysql client
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
		return errors.Wrap(err, "failed to open db connection")
	}
	db.LogMode(true)
	db.SetLogger(NewGormLogger(logger))

	c.conn = db
	c.config = config
	c.logger = logger
	return nil
}

func (c *MysqlClient) GetConn() *gorm.DB {
	return c.conn
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

func (c *MysqlClient) AuthUser(conn *gorm.DB, email string, password string)  (*Auth, error) {
	var auth Auth
	if err := conn.Where("email = ?", email).First(&auth).Error; err != nil {
		return nil, errors.Wrap(err, "cannot check auth exists")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(password)) ; err != nil {
		return nil, errors.Wrap(err, "cannot hash string")
	}
	return &auth, nil
}

//CreateAuth upserts an auth entry to keep idempotence for microservice error handling
func (c *MysqlClient) CreateAuth(conn *gorm.DB, auth Auth) error {
	return conn.Where(Auth{
		UserID:    auth.UserID,
	}).Assign(auth).FirstOrCreate(&auth).Error
}

func (c *MysqlClient) DeleteAuth(conn *gorm.DB, userID uint64) error {
	auth := Auth{}
	auth.UserID = userID
	return conn.Delete(&auth).Error
}