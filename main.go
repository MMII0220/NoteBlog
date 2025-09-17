package main

import (
	"fmt"
	"myasd/config"
	_ "myasd/docs"
	"myasd/internal/controller"
	"myasd/internal/migration"
	"myasd/internal/repository"
	"myasd/internal/service"
)

// @title eBlog service
// @contact.name API eBlog
// @contact.url https://test.com/
// @contact.email nekruzrakhimov@icloud.com
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	dbConn, err := config.StartDBConnection()
	if err != nil {
		fmt.Printf("error in migration: %v", err.Error())
		return
	}

	err = migration.StartMigration(dbConn)
	if err != nil {
		fmt.Printf("error in migration: %v", err.Error())
		return
	}

	repo := repository.NewRepository(dbConn)
	svs := service.NewService(repo)
	ctrl := controller.NewController(svs)

	err = ctrl.StartRoute()
	if err != nil {
		fmt.Printf("error in starting routes: %v", err.Error())
		return
	}

	err = config.CloseDB(dbConn)
	if err != nil {
		fmt.Printf("error in migration: %v", err.Error())
		return
	}
}
