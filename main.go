// main.go
package main

import (
	"MOBILEBANKINGAPI/app"
	"MOBILEBANKINGAPI/router"
	"MOBILEBANKINGAPI/SQLdatabase"
)

func main() {
	err := SQLdatabase.InitDB()
	if err != nil {
		panic("Error initializing database: " + err.Error())
	}

	app.Init()
	r := router.InitRouter()
	r.Run(":8080")
}
