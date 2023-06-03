package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lathief/crm-service/payload/response"
	"github.com/lathief/crm-service/utils/security"
	"net/http"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := security.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.HandleFailedResponse(err.Error(), http.StatusUnauthorized))
			return
		}
		c.Set("adminData", verifyToken)
		c.Next()
	}
}
