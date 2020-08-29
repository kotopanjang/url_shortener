package main

import (
	"aqilliz_assesment/helper"
	"aqilliz_assesment/routes"
	"fmt"
	"os"
)

var (
	err  error
	port = "3112"
)

func main() {

	// if err != nil {
	// 	fmt.Println("statuse: ", err)
	// }

	// defer Config.DB.Close()
	// Config.DB.AutoMigrate(&Models.Todo{})
	err := helper.CheckDatabaseConnection()
	fmt.Println(err)
	if err != nil {
		os.Exit(3)
	}
	r := routes.SetupRouter()
	// running
	helper.WriteLog.Println("Starting App on port ", port)
	r.Run(":" + port)
}
