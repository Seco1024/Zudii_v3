package db

import (
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

func connectRedis() *redis.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	REDIS_HOST := os.Getenv("REDIS_HOST")
	REDIS_PORT := os.Getenv("REDIS_PORT")
	client := redis.NewClient(&redis.Options{
		Addr: REDIS_HOST + ":" + REDIS_PORT,
	})

	_, err = client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GetValueFromRedis(key string) (string, error) {
	client := connectRedis()
	value, err := client.Get(key).Result()
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		log.Fatal(err)
	}
	return value, err
}
func SetValueToRedis(key string, value string, time time.Duration) error {
	client := connectRedis()
	err := client.Set(key, value, time).Err()
	if err != nil {
		log.Fatal(err)
	}
	return err
}
