package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/lathief/crm-service/payload/response"
	"github.com/lathief/crm-service/repository"
	"github.com/lathief/crm-service/utils/database"
	"net/http"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := database.StartDB()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.HandleFailedResponse("Database Error "+err.Error(), http.StatusInternalServerError))
			return
		}
		adminData := c.MustGet("adminData").(jwt.MapClaims)
		adminID := uint(adminData["id"].(float64))
		ActorRepo := repository.ActorNewRepo(db)
		_, err = ActorRepo.GetActorById(adminID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, response.HandleFailedResponse("Account doesn't exist", http.StatusNotFound))
			return
		}
		c.Next()
	}
}
