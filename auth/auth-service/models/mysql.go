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
	return nil
}

func (c *MysqlClient) Migrate() {
	if c.config.AutoMigrate {
		c.conn.AutoMigrate(&User{})
	}
}

func (c *MysqlClient) Close() error {
	err := c.conn.Close()
	return err
}

func (c *MysqlClient) GetUsers() []User {
	var users []User
	c.conn.Limit(10).Find(&users)
	return users
}

func (c *MysqlClient) InsertUser(u User) error {
	var err error
	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(u.Password), 0); err != nil {
		return err
	}

	u.Password = string(hash)

	if err = c.conn.Create(&u).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
