package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name" gorm:"type:varchar(50)"`
	LastName  string `json:"last_name" gorm:"type:varchar(50)"`
	Email     string `json:"email" gorm:"type:varchar(100);unique_index"`
	Password  string `json:"password" gorm:"type:varchar(255)"`
}

func (u User) String() string {
	return fmt.Sprintf("User Id: %i, first name: %s, last name: %s, email: %s", u.ID, u.FirstName, u.LastName, u.Email)
}
