package main

import (
	// "fmt"
	"os"
	"time"

	"myasd/config"
	_ "myasd/docs"
	"myasd/internal/controller"
	"myasd/internal/migration"
	"myasd/internal/repository"
	"myasd/internal/service"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// @title eBlog service
// @contact.name API eBlog
// @contact.url https://test.com/
// @contact.email nekruzrakhimov@icloud.com
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	dbConn, err := config.StartDBConnection()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to DB")
		// fmt.Printf("error in migration: %v", err.Error())
		// return
	}

	err = migration.StartMigration(dbConn)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to run migrations")
		// fmt.Printf("error in migration: %v", err.Error())
		// return
	}

	repo := repository.NewRepository(dbConn, log.Logger)
	svs := service.NewService(repo, log.Logger)
	ctrl := controller.NewController(svs, log.Logger)

	err = ctrl.StartRoute()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start routes")
		// fmt.Printf("error in starting routes: %v", err.Error())
		// return
	}

	err = config.CloseDB(dbConn)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to close DB connection")
		// fmt.Printf("error in migration: %v", err.Error())
		// return
	}
}
