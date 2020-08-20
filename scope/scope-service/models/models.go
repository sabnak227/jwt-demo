package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type UserRole struct {
	gorm.Model
	UserID	  uint64 `json:"user_id" gorm:"type:int(10);unique_index"`
	Role      Role
	RoleID uint
}

func (u UserRole) String() string {
	return fmt.Sprintf("User Role Id: %d, user id: %d, role: %s", u.ID, u.UserID, u.Role.Name)
}

type Role struct {
	gorm.Model
	Name	  string `json:"name" gorm:"type:varchar(50);unique_index"`
	Description string `json:"description" gorm:"type:varchar(255);default:null"`
	Active bool `json:"active" gorm:"type:TINYINT(1);default:1"`
	Permissions         []*Permission `gorm:"many2many:role_permissions;"`
	UserRoles []UserRole
}

func (r Role) String() string {
	return fmt.Sprintf("Role Id: %d, name: %s, description: %s, active: %t", r.ID, r.Name, r.Description, r.Active)
}

type Permission struct {
	gorm.Model
	Name	  string `json:"name" gorm:"type:varchar(50);unique_index"`
	Description string `json:"description" gorm:"type:varchar(255);default:null"`
	Active bool `json:"active" gorm:"type:TINYINT(1);default:1"`
	Roles         []*Role `gorm:"many2many:role_permissions;"`
}

func (p Permission) String() string {
	return fmt.Sprintf("Role Id: %d, name: %s, description: %s, active: %t", p.ID, p.Name, p.Description, p.Active)
}
