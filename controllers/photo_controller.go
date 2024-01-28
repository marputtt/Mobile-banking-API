// controllers/photo_controller.go
package controllers

import (
	
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"MOBILEBANKINGAPI/SQLdatabase"
	"MOBILEBANKINGAPI/helpers"
)

// UploadPhoto handles photo uploading
func UploadPhoto(c *gin.Context) {
	userID, err := helpers.GetUserIDFromToken(c)
	if err != nil {
		helpers.RespondWithError(c, 401, "Unauthorized")
		return
	}

	// Get the file from the request
	file, err := c.FormFile("photo")
	if err != nil {
		helpers.RespondWithError(c, 400, "Bad Request: Missing photo file")
		return
	}

	// Generate a unique filename
	filename := helpers.GenerateUniqueFilename(file.Filename)

	// Save the file to the server
	err = c.SaveUploadedFile(file, filepath.Join("uploads", filename))
	if err != nil {
		helpers.RespondWithError(c, 500, "Error saving photo")
		return
	}

	// Save photo information to the database
	_, err = SQLdatabase.DB.Exec("INSERT INTO photos (user_id, filename) VALUES (?, ?)", userID, filename)
	if err != nil {
		// If there's an error saving to the database, delete the uploaded file
		os.Remove(filepath.Join("uploads", filename))
		helpers.RespondWithError(c, 500, "Error saving photo information")
		return
	}

	helpers.RespondWithJSON(c, http.StatusCreated, gin.H{"message": "Photo uploaded successfully"})
}

// RetrievePhotos handles retrieving user's photos
func RetrievePhotos(c *gin.Context) {
	userID, err := helpers.GetUserIDFromToken(c)
	if err != nil {
		helpers.RespondWithError(c, 401, "Unauthorized")
		return
	}

	// Retrieve user's photos from the database
	rows, err := SQLdatabase.DB.Query("SELECT id, filename FROM photos WHERE user_id = ?", userID)
	if err != nil {
		helpers.RespondWithError(c, 500, "Error retrieving photos")
		return
	}
	defer rows.Close()

	var photos []gin.H
	for rows.Next() {
		var photoID int
		var filename string
		err := rows.Scan(&photoID, &filename)
		if err != nil {
			helpers.RespondWithError(c, 500, "Error scanning photos")
			return
		}
		photos = append(photos, gin.H{"id": photoID, "filename": filename})
	}

	helpers.RespondWithJSON(c, http.StatusOK, gin.H{"photos": photos})
}

// UpdatePhoto handles updating photo information
// UpdatePhoto handles updating photo information
func UpdatePhoto(c *gin.Context) {
	userID, err := helpers.GetUserIDFromToken(c)
	if err != nil {
		helpers.RespondWithError(c, 401, "Unauthorized")
		return
	}

	photoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.RespondWithError(c, 400, "Bad Request: Invalid photo ID")
		return
	}

	// Check if the photo belongs to the user
	var userCheck int
	err = SQLdatabase.DB.QueryRow("SELECT COUNT(*) FROM photos WHERE id = ? AND user_id = ?", photoID, userID).Scan(&userCheck)
	if err != nil || userCheck == 0 {
		helpers.RespondWithError(c, 403, "Forbidden: Photo does not belong to the user")
		return
	}

	var updatedInfo struct {
		Filename string `json:"filename"`
		// Add other fields you want to update
	}

	if err := c.ShouldBindJSON(&updatedInfo); err != nil {
		helpers.RespondWithError(c, 400, "Invalid request")
		return
	}

	// Update photo information in the database
	_, err = SQLdatabase.DB.Exec("UPDATE photos SET filename = ? WHERE id = ?", updatedInfo.Filename, photoID)
	if err != nil {
		helpers.RespondWithError(c, 500, "Error updating photo information")
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, gin.H{"message": "Photo information updated successfully"})
}

// DeletePhoto handles deleting user's photos
func DeletePhoto(c *gin.Context) {
	userID, err := helpers.GetUserIDFromToken(c)
	if err != nil {
		helpers.RespondWithError(c, 401, "Unauthorized")
		return
	}

	photoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.RespondWithError(c, 400, "Bad Request: Invalid photo ID")
		return
	}

	// Check if the photo belongs to the user
	var userCheck int
	err = SQLdatabase.DB.QueryRow("SELECT COUNT(*) FROM photos WHERE id = ? AND user_id = ?", photoID, userID).Scan(&userCheck)
	if err != nil || userCheck == 0 {
		helpers.RespondWithError(c, 403, "Forbidden: Photo does not belong to the user")
		return
	}

	// Retrieve the photo filename from the database
	var filename string
	err = SQLdatabase.DB.QueryRow("SELECT filename FROM photos WHERE id = ?", photoID).Scan(&filename)
	if err != nil {
		helpers.RespondWithError(c, 500, "Error retrieving photo information")
		return
	}

	// Delete the photo record from the database
	_, err = SQLdatabase.DB.Exec("DELETE FROM photos WHERE id = ?", photoID)
	if err != nil {
		helpers.RespondWithError(c, 500, "Error deleting photo information")
		return
	}

	// Delete the photo file from the server
	err = os.Remove(filepath.Join("uploads", filename))
	if err != nil {
		helpers.RespondWithError(c, 500, "Error deleting photo file")
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}
