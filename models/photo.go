// models/photo.go
package models

import (
	"MOBILEBANKINGAPI/SQLdatabase"
	
)

// Photo represents the model for a photo
type Photo struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Filename string `json:"filename"`
	// Add other fields as needed
}

// InitPhotos initializes the photo model in the database
func InitPhotos() {
	// Execute SQL statements to create the 'photos' table if it doesn't exist
	_, err := SQLdatabase.DB.Exec(`
		CREATE TABLE IF NOT EXISTS photos (
			id INT AUTO_INCREMENT PRIMARY KEY,
			user_id INT,
			filename VARCHAR(255)
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		)
	`)
	if err != nil {
		panic(err.Error())
	}

	// You can add more initialization logic here if needed
}

// init initializes the photo model

