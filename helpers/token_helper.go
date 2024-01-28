// helpers/token_helper.go
package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"errors"
)

// Error constants for token handling
var (
	ErrNoAuthToken       = errors.New("no authorization token provided")
	ErrInvalidAuthToken  = errors.New("invalid authorization token")
)

// GetUserIDFromToken retrieves the user ID from the JWT token in the request header
func GetUserIDFromToken(c *gin.Context) (int64, error) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return 0, ErrNoAuthToken
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("0p6FMYEeSj"), nil
	})
	if err != nil || !token.Valid {
		return 0, ErrInvalidAuthToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, ErrInvalidAuthToken
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, ErrInvalidAuthToken
	}

	return int64(userID), nil
}
