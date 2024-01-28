// middlewares/auth_middleware.go
package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"MOBILEBANKINGAPI/helpers"
)

// AuthMiddleware is a middleware to handle JWT authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			helpers.RespondWithError(c, http.StatusUnauthorized, "Unauthorized: Missing Authorization header")
			c.Abort()
			return
		}

		// Check if the header format is valid (Bearer token)
		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 {
			helpers.RespondWithError(c, http.StatusUnauthorized, "Unauthorized: Invalid Authorization header format")
			c.Abort()
			return
		}

		tokenString := splitToken[1]

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check the signing method and provide the secret key
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			//secret key
			return []byte("0p6FMYEeSj"), nil
		})

		if err != nil {
			helpers.RespondWithError(c, http.StatusUnauthorized, "Unauthorized: Invalid token")
			c.Abort()
			return
		}

		// Check if the token is valid
		if !token.Valid {
			helpers.RespondWithError(c, http.StatusUnauthorized, "Unauthorized: Token is not valid")
			c.Abort()
			return
		}

		// Extract user ID from the token claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			helpers.RespondWithError(c, http.StatusUnauthorized, "Unauthorized: Unable to extract user ID from token")
			c.Abort()
			return
		}

		userID, ok := claims["user_id"].(float64)
		if !ok {
			helpers.RespondWithError(c, http.StatusUnauthorized, "Unauthorized: Invalid user ID in token")
			c.Abort()
			return
		}

		// Set the user ID in the request context for further processing
		c.Set("user_id", int64(userID))
		c.Next()
	}
}
