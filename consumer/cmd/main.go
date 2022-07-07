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
	timeToConnectDB    = 5
	attemptsDB         = 5
	timeToConnectRedis = 5
	attemptsRedis      = 5
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	logr := logger.GetLogger(cfg.Logger.LogLevel)

	postgresClient, err := postgres.NewClient(cfg, attemptsDB, timeToConnectDB)
	if err != nil {
		logr.Fatal(err)
	}
	logr.Trace("connected to DB")

	fmt.Println(postgresClient)
}
