package main

import (
	"api-key-value-service/redis"
	"api-key-value-service/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Println("Starting appliction....")
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
	config := AppConfig{}

	config.LoadEnv()
	redis.CreateRedisClient()
	log.Println("Intializing server")
	server.InitializeServer()
}
