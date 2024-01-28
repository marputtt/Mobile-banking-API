// app/app.go
package app

import (
	"MOBILEBANKINGAPI/SQLdatabase"
	"MOBILEBANKINGAPI/models"
)

func Init() {
	SQLdatabase.InitDB()

	// Initialize models
	models.InitModels()
	

}
