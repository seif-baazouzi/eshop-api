package db

import (
	"fmt"
	"os"

	"github.com/gomodule/redigo/redis"
)

var Redis *redis.Pool

func InitRedis() {
	dbHostname := os.Getenv("REDIS_HOSTNAME")

	Redis = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", dbHostname+":6379")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Unable to connect to redis: %v\n", err)
				os.Exit(1)
			}

			return c, err
		},
	}
}
