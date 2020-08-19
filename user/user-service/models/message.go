package models

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	UserMsgTypeCreated = "created"
	UserMsgTypeDeleted = "deleted"
)

type UserMsg struct {
	Type      string `json:"type"`
	Status    string `json:"status"`
	UserId    uint64 `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func(u *UserMsg) SetPassword(password string) *UserMsg{
	// hashing password
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u.Password =string(hash)
	return u
}

func NewUserMsg(user *User, msgType string) *UserMsg{
	return &UserMsg{
		Type:      msgType,
		Status:    user.Status,
		UserId:    uint64(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}