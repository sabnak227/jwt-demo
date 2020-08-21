package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

const (
	UserStatusEnabled = "enable"
	UserStatusDisabled = "disable"
	UserStatusPending = "pending"
)

type User struct {
	gorm.Model
	Status    string `json:"status" gorm:"type:varchar(50)"`
	FirstName string `json:"first_name" gorm:"type:varchar(50)"`
	LastName  string `json:"last_name" gorm:"type:varchar(50)"`
	Email     string `json:"email" gorm:"type:varchar(100);unique_index"`
	Address1  string `json:"address1" gorm:"type:varchar(255);default:null"`
	Address2  string `json:"address2" gorm:"type:varchar(255);default:null"`
	City      string `json:"city" gorm:"type:varchar(30);default:null"`
	State     string `json:"state" gorm:"type:varchar(30);default:null"`
	Country   string `json:"country" gorm:"type:varchar(30);default:null"`
	Phone     string `json:"phone" gorm:"type:varchar(50);default:null"`
}

func (u User) String() string {
	return fmt.Sprintf("User Id: %i, status %s, first name: %s, last name: %s, email: %s", u.ID, u.Status, u.FirstName, u.LastName, u.Email)
}
