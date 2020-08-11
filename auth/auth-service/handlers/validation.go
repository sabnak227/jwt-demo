package handlers

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	pb "github.com/sabnak227/jwt-demo/auth"
)


type loginRequest struct {
	req pb.LoginRequest
}

func (l loginRequest) Validate() error {
	s := l.req
	return validation.ValidateStruct(&s,
		validation.Field(&s.Email, is.Email),
		validation.Field(&s.Password, validation.Required, validation.Length(6,50)),
	)
}

type createAuthRequest struct {
	req pb.CreateAuthRequest
}

func (l createAuthRequest) Validate() error {
	s := l.req
	return validation.ValidateStruct(&s,
		validation.Field(&s.UserId, validation.Required),
		validation.Field(&s.Email, validation.Required, is.Email),
		validation.Field(&s.Password, validation.Required, validation.Length(6,50)),
		validation.Field(&s.FirstName, validation.Required, validation.Length(3, 50)),
		validation.Field(&s.LastName, validation.Required, validation.Length(3,50)),
	)
}

type refreshRequest struct {
	req pb.RefreshRequest
}

func (l refreshRequest) Validate() error {
	s := l.req
	return validation.ValidateStruct(&s,
		validation.Field(&s.RefreshToken, validation.Required),
	)
}