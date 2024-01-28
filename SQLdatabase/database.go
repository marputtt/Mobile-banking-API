// SQLdatabase/database.go
package SQLdatabase

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// InitializeDB initializes the database connection
func InitDB() error {
	// Replace the connection parameters with your actual database details
	dsn := "marput:1234@tcp(localhost:3306)/mobile_banking"

	// Open a connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// Check if the connection is successful
	if err := db.Ping(); err != nil {
		return err
	}

	// Set the global DB variable
	DB = db

	fmt.Println("Connected to the database")

	return nil
}
