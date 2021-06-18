package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"os"
)

var redisClient redis.Client

//CreateRedisClient -
func CreateRedisClient() {
	redisURL := os.Getenv("REDIS_URL")

	cl := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := cl.Ping(context.TODO()).Result()
	if err != nil {
		log.Errorln("Redis Connection Failed")
		log.Fatalln(err.Error())
	}
	log.Infoln("Redis Connected !")
	redisClient = *cl
}

//GetRedisClient -
func GetRedisClient() *redis.Client {
	return &redisClient
}
