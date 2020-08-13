package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Auth struct {
	gorm.Model
	UserID	  uint64 `json:"user_id" gorm:"type:int(10);unique_index"`
	FirstName string `json:"first_name" gorm:"type:varchar(50)"`
	LastName  string `json:"last_name" gorm:"type:varchar(50)"`
	Email     string `json:"email" gorm:"type:varchar(100);unique_index"`
	Password  string `json:"password" gorm:"type:varchar(255)"`
}

func (a Auth) String() string {
	return fmt.Sprintf("User Id: %i, first name: %s, last name: %s, email: %s", a.ID, a.FirstName, a.LastName, a.Email)
}
