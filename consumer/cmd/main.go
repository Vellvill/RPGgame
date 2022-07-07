package main

import (
	cache "Consumer/pkg/redis"
	standartLog "log"

	log "github.com/sirupsen/logrus"

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
		standartLog.Fatal(err)
	}

	logr := logger.GetLogger(cfg.Logger.LogLevel)

	postgresClient, err := postgres.NewPGClient(cfg, attemptsDB, timeToConnectDB)
	if err != nil {
		logr.WithFields(log.Fields{
			"type": "postgres",
		}).Fatal(err)
	}

	logr.WithFields(log.Fields{
		"type": "postgres",
	}).Info("connected to postgres")

	redisClient, err := cache.NewRedisClient(cfg, timeToConnectRedis, attemptsRedis)
	if err != nil {
		logr.WithFields(log.Fields{
			"type": "redis",
		}).Fatal(err)
	}

	logr.WithFields(log.Fields{
		"type": "redis",
	}).Info("connected to redis")

}
