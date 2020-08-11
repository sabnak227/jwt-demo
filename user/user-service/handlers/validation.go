package handlers

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	pb "github.com/sabnak227/jwt-demo/user"
)


type createUserRequest struct {
	req pb.CreateUserRequest
}

func (l createUserRequest) Validate() error {
	s := l.req
	return validation.ValidateStruct(&s,
		validation.Field(&s.FirstName, validation.Required, validation.Length(3, 50)),
		validation.Field(&s.LastName, validation.Required, validation.Length(3,50)),
		validation.Field(&s.Email, validation.Required, is.Email),
		validation.Field(&s.Password, validation.Required, validation.Length(6,50)),
		validation.Field(&s.Address1, validation.Length(0,255)),
		validation.Field(&s.Address2, validation.Length(0,255)),
		validation.Field(&s.City, validation.Length(0,30)),
		validation.Field(&s.State, validation.Length(0,30)),
		validation.Field(&s.Country, validation.Length(0,30)),
		validation.Field(&s.Phone, validation.Length(0,50)),
	)
}
