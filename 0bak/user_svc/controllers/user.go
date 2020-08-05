package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sabnak/jwt-demo/user_svc/models"
)

func (ctl *Controller) GetUsers(c *gin.Context) {
	ctl.logger.Log("msg", "Getting all users")
	u := ctl.dbClient.GetUsers()

	c.JSON(http.StatusOK, gin.H{"data": u})
}

type InsertUserInput struct {
	FirstName string `form:"first_name" xml:"first_name" json:"first_name" binding:"required"`
	LastName  string `form:"last_name" xml:"last_name" json:"last_name" binding:"required"`
	// Email     string `form:"email" xml:"email" json:"email" binding:"required,email,unique=users"`
	Email    string `form:"email" xml:"email" json:"email" binding:"required,email"`
	Password string `form:"password" xml:"password" json:"password" binding:"required"`
	// PasswordConfirmation string `form:"password_confirmation" xml:"password_confirmation" json:"password_confirmation" binding:"required" validate:"required,eqfield=Password"`
	Address1 string `form:"address1" xml:"address1" json:"address1"`
	Address2 string `form:"address2" xml:"address2" json:"address2"`
	City     string `form:"city" xml:"city" json:"city"`
	State    string `form:"state" xml:"state" json:"state"`
	Country  string `form:"country" xml:"country" json:"country"`
	Phone    string `form:"phone" xml:"phone" json:"phone"`
}

func (ctl *Controller) InsertUser(c *gin.Context) {
	var input InsertUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
		Address1:  input.Address1,
		Address2:  input.Address2,
		City:      input.City,
		State:     input.State,
		Country:   input.Country,
		Phone:     input.Phone,
	}
	ctl.logger.Log("msg", "Inserting new users", "user", u.String())

	if err := ctl.dbClient.InsertUser(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": u})
}
