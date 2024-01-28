// helpers/response_helper.go
package helpers

import (
	"github.com/gin-gonic/gin"
	
)

// RespondWithError sends a JSON error response with the specified status code and message
func RespondWithError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message})
}

// RespondWithJSON sends a JSON response with the specified status code and data
func RespondWithJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}
