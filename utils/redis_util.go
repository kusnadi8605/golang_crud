package util

import (
	"fmt"
	conf "promo_api/config"

	"github.com/gomodule/redigo/redis"
)

//Ping ..
func Ping() error {
	conn := conf.RedisPool.Get()
	defer conn.Close()

	r, err := redis.String(conn.Do("PING"))
	if err != nil {
		return fmt.Errorf("cannot 'PING' db: %v", err)
	}

	fmt.Println(r)
	return nil
}

//SetTex  ..
func SetTex(key string, time int, value []byte) error {

	conn := conf.RedisPool.Get()
	defer conn.Close()

	_, err := conn.Do("SETEX", key, time, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return fmt.Errorf("error setting key %s to %s: %v", key, v, err)
	}
	return err
}

//Get ..
func Get(key string) ([]byte, error) {
	conn := conf.RedisPool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	/*
		if err != nil {
			return data, fmt.Errorf("error getting key %s: %v", key, err)
		}*/
	return data, err
}

//Set ..
func Set(key string, value []byte) error {

	conn := conf.RedisPool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return fmt.Errorf("error setting key %s to %s: %v", key, v, err)
	}
	return err
}

//Exists .
func Exists(key string) (bool, error) {

	conn := conf.RedisPool.Get()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return ok, fmt.Errorf("error checking if key %s exists: %v", key, err)
	}
	return ok, err
}

//Delete .
func Delete(key string) error {

	conn := conf.RedisPool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	return err
}

//GetKeys .
func GetKeys(pattern string) ([]string, error) {

	conn := conf.RedisPool.Get()
	defer conn.Close()

	iter := 0
	keys := []string{}
	for {
		arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", pattern))
		if err != nil {
			return keys, fmt.Errorf("error retrieving '%s' keys", pattern)
		}

		iter, _ = redis.Int(arr[0], nil)
		k, _ := redis.Strings(arr[1], nil)
		keys = append(keys, k...)

		if iter == 0 {
			break
		}
	}

	return keys, nil
}

//Incr ..
func Incr(counterKey string) (int, error) {

	conn := conf.RedisPool.Get()
	defer conn.Close()

	return redis.Int(conn.Do("INCR", counterKey))
}
