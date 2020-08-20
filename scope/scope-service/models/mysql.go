package models

import (
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sabnak227/jwt-demo/scope/scope-service/config"
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

func (c *MysqlClient) GetConn() *gorm.DB {
	return c.conn
}

func (c *MysqlClient) Migrate() {
	if c.config.AutoMigrate {
		c.conn.AutoMigrate(&Role{})
		c.conn.AutoMigrate(&Permission{})
		c.conn.AutoMigrate(&UserRole{}).AddForeignKey("role_id", "roles(id)", "CASCADE", "CASCADE")
		if c.config.Seed {
			c.Seed()
		}
	}
}

func (c *MysqlClient) Close() error {
	err := c.conn.Close()
	return err
}

func (c *MysqlClient) Seed() {
	var role Role
	var perms []string
	conn := c.conn.Begin()
	role = Role{
		Name: "admin",
		Description: "admin user role",
	}

	perms = []string{
		"read playlog",
		"write playlog",
		"read comment",
		"write comment",
	}

	if err := c.seedRollPerms(conn, role, perms); err != nil {
		conn.Rollback()
	}

	role = Role{
		Name: "user",
		Description: "normal user role",
	}

	perms = []string{
		"read comment",
		"write comment",
	}

	if err := c.seedRollPerms(conn, role, perms); err != nil {
		conn.Rollback()
	}

	conn.Commit()
}

func (c *MysqlClient) seedRollPerms(conn *gorm.DB, role Role, perms []string) error{
	roleDb, err := c.CreateRole(conn, role)
	if err != nil {
		c.logger.Fatalf("Failed to create role, err: %s", err)
		return err
	}

	for i := 0; i < len(perms); i++  {
		perm := Permission{
			Name:        perms[i],
		}
		permDb, err := c.CreatePermission(conn, perm)
		if err != nil {
			c.logger.Fatalf("Failed to create permission, err: %s", err)
			return err
		}
		if err := c.AttachPerm(conn, *roleDb, *permDb); err != nil {
			c.logger.Fatalf("Failed to attach permission to role, err: %s", err)
			return err
		}
	}
	return nil
}

func (c *MysqlClient) GetPerms(conn *gorm.DB, userID uint64) ([]string, error) {
	var userRole UserRole
	if err := conn.Preload("Role.Permissions").
		Where("user_id = ?", userID).
		Find(&userRole).Error; err != nil {
		return []string{}, err
	}

	perms := userRole.Role.Permissions
	var res []string
	for i:=0;i < len(perms);i++  {
		res = append(res, perms[i].Name)
	}

	return res, nil
}

// CreateRole upserts a role entry
func (c *MysqlClient) CreateRole(conn *gorm.DB, role Role) (*Role, error) {
	var res Role
	db := conn.Where(Role{
		Name:    role.Name,
	}).Assign(role).FirstOrCreate(&res)
	return &res, db.Error
}

// CreatePermission upserts a permission entry
func (c *MysqlClient) CreatePermission(conn *gorm.DB, perm Permission) (*Permission, error) {
	var res Permission
	db := conn.Where(Permission{
		Name:    perm.Name,
	}).Assign(perm).FirstOrCreate(&res)
	return &res, db.Error
}

func (c *MysqlClient) AttachPerm(conn *gorm.DB, role Role, perm Permission) error {
	return conn.Model(&perm).Association("Roles").Append(role).Error
}

func (c *MysqlClient) DetachPerm(conn *gorm.DB, role Role, perm Permission) error {
	return conn.Model(&perm).Association("Roles").Delete(role).Error
}

func (c *MysqlClient) AssignRole(conn *gorm.DB, userID uint64, roleName string) error {
	var role Role
	if err := conn.Where("name = ?", roleName).First(&role).Error; err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to get role %s", roleName))
	}
	// upsert user role
	var res UserRole
	if err := conn.Where(UserRole{
		UserID: userID,
		RoleID: role.ID,
	}).FirstOrCreate(&res).Error; err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to upser user role"))
	}

	return nil
}