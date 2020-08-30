package main

import (
	"os"

	"github.com/kotopanjang/aqilliz_assesment/helper"
	"github.com/kotopanjang/aqilliz_assesment/routes"
)

var (
	err error
)

func main() {
	port := os.Getenv("appport")
	if port == "" {
		helper.Println("env 'port' Required!")
		os.Exit(3)
	}
	dbhost := os.Getenv("dbhost")
	if dbhost == "" {
		helper.Println("env 'dbhost' Required!")
		os.Exit(3)
	}
	dbport := os.Getenv("dbport")
	if dbport == "" {
		helper.Println("env 'dbport' Required!")
		os.Exit(3)
	}
	db := os.Getenv("db")
	if db == "" {
		helper.Println("env 'db' Required!")
		os.Exit(3)
	}
	dbuser := os.Getenv("dbuser")
	dbpass := os.Getenv("dbpass")

	helper.Println("Starting App on port ", port)

	helper.WriteConfig(dbhost, dbport, dbuser, dbpass, db)
	if err != nil {
		helper.Println(err)
		os.Exit(3)
	}

	err := helper.CheckDatabaseConnection()
	if err != nil {
		helper.Println(err.Error())
		os.Exit(3)
	}
	r := routes.SetupRouter()
	// running
	helper.Println("App Started on port ", port)
	r.Run(":" + port)
}
