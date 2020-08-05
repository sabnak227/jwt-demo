package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (ctl *Controller) Login(c *gin.Context) {
	ctl.logger.Log("msg", "Logging in user")

	user := "jason"

	// create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("RS256"))
	claims := make(jwt.MapClaims)

	// set our claims
	claims["AccessToken"] = "level1"
	claims["CustomUserInfo"] = struct {
		Name string
		Kind string
	}{user, "human"}

	// set the expire time
	// see http://tools.ietf.org/html/draft-ietf-oauth-json-web-token-20#section-4.1.4
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	t.Claims = claims

	tokenString, err := t.SignedString(ctl.signKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to sign token")
		ctl.logger.Log("error", "Failed to sign token")
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenString})
}
