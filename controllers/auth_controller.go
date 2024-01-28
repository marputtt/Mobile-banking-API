// controllers/auth_controller.go
package controllers

import (
	"time"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"MOBILEBANKINGAPI/SQLdatabase"
	"MOBILEBANKINGAPI/helpers"
	"MOBILEBANKINGAPI/models"
)

// RegisterUser handles user registration
func RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.RespondWithError(c, 400, "Invalid request")
		return
	}

	// Hash the password before saving it to the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.RespondWithError(c, 500, "Error hashing password")
		return
	}

	user.Password = string(hashedPassword)

	// Save the user to the database
	if err := SQLdatabase.DB.QueryRow("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password).Scan(&user.ID); err != nil {
		helpers.RespondWithError(c, 500, "Error creating user")
		return
	}

	user.Password = "" // Do not send the password in the response
	helpers.RespondWithJSON(c, 201, user)
}

// LoginUser handles user login
func LoginUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.RespondWithError(c, 400, "Invalid request")
		return
	}

	// Retrieve the user from the database
	err := SQLdatabase.DB.QueryRow("SELECT id, password FROM users WHERE username = ?", user.Username).Scan(&user.ID, &user.Password)
	if err != nil {
		helpers.RespondWithError(c, 401, "Invalid credentials")
		return
	}

	// Compare the stored hashed password with the entered password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password))
	if err != nil {
		helpers.RespondWithError(c, 401, "Invalid credentials")
		return
	}

	// You may generate and send a JWT token here for authentication
	// ...
	 // Generate JWT token
	 token, err := generateToken(user.ID)
	 if err != nil {
		 helpers.RespondWithError(c, 500, "Error generating token")
		 return
	 }
 
	

	 user.Password = "" // Do not send the password in the response
	 helpers.RespondWithJSON(c, 200, gin.H{"token": token})
}

// UpdateUser handles updating user information
func UpdateUser(c *gin.Context) {
	// Implement logic to update user information
	helpers.RespondWithError(c, 501, "Not Implemented")
}

// DeleteUser handles deleting user accounts
func DeleteUser(c *gin.Context) {
	// Implement logic to delete user accounts
	helpers.RespondWithError(c, 501, "Not Implemented")
}

func generateToken(userID int64) (string, error) {
    // Define the JWT claims
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (e.g., 24 hours)
    }

    // Create the JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign the token with a secret key
    signedToken, err := token.SignedString([]byte("0p6FMYEeSj"))
    if err != nil {
        return "", err
    }

    return signedToken, nil
}