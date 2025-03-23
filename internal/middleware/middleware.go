package middleware

import (
	"net/http"
	"os"
	"strings"
	"tasks/config"
	"tasks/internal/models"
	"tasks/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

/*
-----------------------------------------------------------------------------------------

	Function Name : AuthMiddleware
	Purpose       : Validate the JWT token and authorize API requests.

-----------------------------------------------------------------------------------------

		Returns:
	   -----------------------------------------------------------------------------------------
		 gin.HandlerFunc

		Success Response:
		-----------------------------------------------------------------------------------------
		If successful, the function allows the request to proceed with userID in the context.

		Error Handling:
		-----------------------------------------------------------------------------------------
		On failure, an unauthorized status is returned, and the request is aborted.

		Author        : LOGESHKUMAR P
		Created Date  : 22-03-2025

-----------------------------------------------------------------------------------------
*/
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log := new(utils.Logger)
		log.SetSid(ctx.Request)
		log.Log(utils.INFO, "AuthMiddleware start")
		defer log.Log(utils.INFO, "AuthMiddleware end")

		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			log.Log(utils.ERROR, "AM001", "Missing or invalid Authorization header")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			log.Log(utils.ERROR, "AM002", "Invalid or expired token: "+err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}
		if !token.Valid {
			log.Log(utils.ERROR, "AM002", "Invalid or expired token: ")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		if claims.UserID == 0 {
			log.Log(utils.ERROR, "AM003", "Missing user ID in token payload")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token payload"})
			return
		}

		ctx.Set("userID", claims.UserID)

		ctx.Next()
	}
}

func IPRestrictionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientIP := ctx.ClientIP()
		var cf models.Config

		err := config.LoadTOML("toml/config.toml", &cf)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access restricted"})
			return
		}
		for _, ip := range cf.AllowedIPs {
			if clientIP == ip {
				ctx.Next()
				return
			}
		}

		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access restricted"})
	}
}
