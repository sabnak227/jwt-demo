package models

import (
	"fmt"
	"github.com/pkg/errors"
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
		c.conn.AutoMigrate(&User{})
	}
}

func (c *MysqlClient) Close() error {
	err := c.conn.Close()
	return err
}

func (c *MysqlClient) ListUser(conn *gorm.DB, offset uint32, limit uint32, search string) (*[]User, error) {
	var users []User
	str := "%" + search + "%"
	err := conn.Offset(offset).Limit(limit).
		Unscoped().
		Where("first_name LIKE ?", str).
		Or("last_name LIKE ?", str).
		Or("email LIKE ?", str).
		Or("status LIKE ?", str).
		Find(&users).Error
	return &users, err
}

func (c *MysqlClient) GetUser(conn *gorm.DB, id uint64) (*User, error) {
	var user User
	err := conn.Where("id = ?", id).Unscoped().Find(&user).Error
	return &user, err
}

func (c *MysqlClient) CheckEmailExists(conn *gorm.DB, email string) (bool, error) {
	var cnt int
	err := conn.Model(&User{}).Unscoped().Where("email = ?", email).Count(&cnt).Error
	return cnt > 0, err
}

func (c *MysqlClient) CreateUser(conn *gorm.DB, user User) (*User, error) {
	if user.Status != "" {
		user.Status = UserStatusEnabled
	}
	err := conn.Create(&user).Error
	return &user, err
}

func (c *MysqlClient) UpdateUser(conn *gorm.DB, user User) (*User, error) {
	var res User
	err := conn.Model(&res).Unscoped().Updates(user).Error
	return &res, err
}

func (c *MysqlClient) Delete(conn *gorm.DB, id uint64) error {
	user := User{}
	user.ID = uint(id)
	return conn.Delete(&user).Error
}