package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lathief/crm-service/payload/response"
	"github.com/lathief/crm-service/repository"
	"github.com/lathief/crm-service/utils/database"
	"net/http"
	"strconv"
)

func AdminAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := database.StartDB()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.HandleFailedResponse("Database Error "+err.Error(), http.StatusInternalServerError))
			return
		}
		adminId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.HandleFailedResponse("Invalid Path Variable "+err.Error(), http.StatusBadRequest))
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
		if uint(adminId) != adminID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.HandleFailedResponse("You are not allowed to access this data", http.StatusBadRequest))
			return
		}
		c.Next()
	}
}
func SuperAdminAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := database.StartDB()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.HandleFailedResponse("Database Error "+err.Error(), http.StatusInternalServerError))
			return
		}
		adminData := c.MustGet("adminData").(jwt.MapClaims)
		adminID := uint(adminData["id"].(float64))
		ActorRepo := repository.ActorNewRepo(db)
		spAdmin, err := ActorRepo.GetActorById(adminID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, response.HandleFailedResponse("Account doesn't exist", http.StatusNotFound))
			return
		}
		if spAdmin.Role.Rolename != "role_super_admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.HandleFailedResponse("You are not allowed to access this data", http.StatusBadRequest))
			return
		}
		c.Next()
	}
}
