package main

import (
	"fmt"
	"myasd/config"
	_ "myasd/docs"
	"myasd/internal/controller"
	"myasd/internal/migration"
)

// @title eBlog service
// @contact.name API eBlog
// @contact.url https://test.com/
// @contact.email nekruzrakhimov@icloud.com
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	err := config.StartDBConnection()
	if err != nil {
		fmt.Printf("error in migration: %v", err.Error())
		return
	}

	err = migration.StartMigration()
	if err != nil {
		fmt.Printf("error in migration: %v", err.Error())
		return
	}

	err = controller.StartRoute()
	if err != nil {
		fmt.Printf("error in starting routes: %v", err.Error())
		return
	}

	err = config.CloseDB()
	if err != nil {
		fmt.Printf("error in migration: %v", err.Error())
		return
	}
}
