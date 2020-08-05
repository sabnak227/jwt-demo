package models

import "fmt"

type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	FirstName string `json:"first_name" gorm:"type:varchar(50)"`
	LastName  string `json:"last_name" gorm:"type:varchar(50)"`
	Email     string `json:"email" gorm:"type:varchar(100);unique_index"`
	Password  string `json:"password" gorm:"type:varchar(255)"`
	Address1  string `json:"address1" gorm:"type:varchar(255)"`
	Address2  string `json:"address2" gorm:"type:varchar(255)"`
	City      string `json:"city" gorm:"type:varchar(30)"`
	State     string `json:"state" gorm:"type:varchar(30)"`
	Country   string `json:"country" gorm:"type:varchar(30)"`
	Phone     string `json:"phone" gorm:"type:varchar(50)"`
}

func (u User) String() string {
	return fmt.Sprintf("User Id: %i", u.ID)
}
