package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lathief/crm-service/utils/security"
)

type Auth struct {
}

type AuthorizationInterface interface {
	Authentication(c *gin.Context) error
}

func NewSecurity() AuthorizationInterface {
	return &Auth{}
}

func (auth *Auth) Authentication(c *gin.Context) error {
	verifyToken, err := security.VerifyToken(c)
	if err != nil {
		return err
	}
	c.Set("adminData", verifyToken)
	return nil
}
