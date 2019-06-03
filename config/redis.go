package config

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var (
	//RedisPool ..
	RedisPool *redis.Pool
)

//RedisDbInit init
func RedisDbInit(redisURL string) {
	RedisPool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisURL)
			if err != nil {
				panic(err.Error())
				//fmt.Println("Error Redis")
			}
			fmt.Println("Initiating redis connection: ", redisURL)
			return c, nil
		}, /*TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},*/
	}
}

//https://medium.com/@gilcrest_65433/basic-redis-examples-with-go-a3348a12878e

//NewPool ..
/*
func NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     10000000,                 // Maximum number of idle connections in the pool.
		MaxActive:   200,                      // Maximum number of connections allocated by the pool at a given time.
		IdleTimeout: 5000000000 * time.Second, // Close connections after remaining idle for this duration. Applications should set the timeout to a value less than the server's timeout.
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		}, TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

}*/
