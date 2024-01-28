// controllers/user_controller.go
package controllers

import (
	"github.com/gin-gonic/gin"
	
	"MOBILEBANKINGAPI/SQLdatabase"
	"MOBILEBANKINGAPI/helpers"
	"MOBILEBANKINGAPI/models"
)

// UpdateUserInfo handles updating user information
func UpdateUserInfo(c *gin.Context) {
	userID, err := helpers.GetUserIDFromToken(c)
	if err != nil {
		helpers.RespondWithError(c, 401, "Unauthorized")
		return
	}

	var updatedUser models.User

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		helpers.RespondWithError(c, 400, "Invalid request")
		return
	}

	// Update user information in the database
	_, err = SQLdatabase.DB.Exec("UPDATE users SET email = ?, full_name = ? WHERE id = ?", updatedUser.Email, updatedUser.FullName, userID)
	if err != nil {
		helpers.RespondWithError(c, 500, "Error updating user information")
		return
	}

	helpers.RespondWithJSON(c, 200, gin.H{"message": "User information updated successfully"})
}

// DeleteUserAccount handles deleting user accounts
func DeleteUserAccount(c *gin.Context) {
	userID, err := helpers.GetUserIDFromToken(c)
	if err != nil {
		helpers.RespondWithError(c, 401, "Unauthorized")
		return
	}

	// Delete user account from the database
	_, err = SQLdatabase.DB.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		helpers.RespondWithError(c, 500, "Error deleting user account")
		return
	}

	helpers.RespondWithJSON(c, 200, gin.H{"message": "User account deleted successfully"})
}
