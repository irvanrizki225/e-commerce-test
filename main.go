package main

import (
	"e-commerce/utilities"
	"e-commerce/routes"
	// "e-commerce/models"

	"log"
	"net/http"
)

var db = utilities.ConnecDB()

func main() {
	// models.MigrateUser(db)
	// models.MigrateProduct(db)

	routes.SetRouter()

	log.Println("Server running at port 9060")
	log.Fatal(http.ListenAndServe(":9060", nil))

	runBankAccountSystem()
}