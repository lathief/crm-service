package security

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lathief/crm-service/config"
	"github.com/lathief/crm-service/constant"
	"log"
	"strings"
	"time"
)

func GenerateToken(id uint, username string) string {
	claims := jwt.MapClaims{
		"id":       id,
		"username": username,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Second * time.Duration(config.Config.AuthKey.ExpiresAt)).Unix(),
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := parseToken.SignedString([]byte(config.Config.AuthKey.SecretKey))
	if err != nil {
		log.Fatal(err)
	}
	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) {

	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, constant.ErrAdminNeedLogin
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, constant.ErrTokenInvalid
		}
		return []byte(config.Config.AuthKey.SecretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, constant.ErrTokenInvalid
	}

	return token.Claims.(jwt.MapClaims), nil
}
