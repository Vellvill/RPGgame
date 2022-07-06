package main

import (
	"fmt"
	"log"

	"Consumer/internal/config"
	"Consumer/pkg/logger"
	"Consumer/pkg/postgres"
)

//parameters for postgres connection
const (
	timeToConnect = 5
	attempts      = 5
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	logr := logger.GetLogger(cfg.Logger.LogLevel)

	postgresClient, err := postgres.NewClient(cfg, attempts, timeToConnect)
	if err != nil {
		logr.Fatal(err)
	}

	fmt.Println(postgresClient)
}
