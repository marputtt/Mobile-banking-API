// models/user.go
package models

import (
	"MOBILEBANKINGAPI/SQLdatabase"
	"database/sql"
	"fmt"
)

// User represents the user model
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

// InitModels initializes user-related models
func InitModels() {
	// Ensure the users table exists in the database
	_, err := SQLdatabase.DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			full_name VARCHAR(255) NOT NULL
		);
	`)
	if err != nil {
		fmt.Println("Error creating users table:", err)
	}
}

// GetUserByID retrieves a user by ID from the database
func GetUserByID(userID int64) (*User, error) {
	var user User
	err := SQLdatabase.DB.QueryRow("SELECT id, username, email, full_name FROM users WHERE id = ?", userID).
		Scan(&user.ID, &user.Username, &user.Email, &user.FullName)
	if err == sql.ErrNoRows {
		return nil, nil // No user found
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}
