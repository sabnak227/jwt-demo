package models

import (
	"fmt"
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sabnak227/jwt-demo/user/user-service/config"
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
		c.conn.AutoMigrate(&User{})
	}
}

func (c *MysqlClient) Close() error {
	err := c.conn.Close()
	return err
}

func (c *MysqlClient) GetUser(id uint64) (*User, error) {
	var user User
	err := c.conn.Where("id = ?", id).Find(&user).Error
	return &user, err
}

func (c *MysqlClient) CheckEmailExists(email string) (*User, error) {
	var user User
	err := c.conn.Where("email = ?", email).Find(&user).Error
	return &user, err
}

func (c *MysqlClient) CreateUser(user User) (*User, error) {
	err := c.conn.Create(&user).Error
	return &user, err
}

func (c *MysqlClient) Delete(id uint64) error {
	user := User{}
	user.ID = uint(id)
	return c.conn.Delete(&user).Error
}