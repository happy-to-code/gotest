package _func

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

// SetXX
func SetEx(key string, value interface{}, time int) bool {
	defer c.Close()
	_, err := c.Do("set", key, value, "EX", time)
	if err != nil {
		log.Println("set a value error,", err)
		return false
	}
	return true
}

// string 操作
func Set(key string, value interface{}) bool {
	defer c.Close()
	_, err := c.Do("set", key, value)
	if err != nil {
		log.Println("set a value error,", err)
		return false
	}
	return true
}

func Get(key string) string {
	defer c.Close()
	value, err := redis.String(c.Do("get", key))
	if err != nil {
		log.Println("get a value error,", err)
		return ""
	}

	return value
}

func Del(key string) bool {
	defer c.Close()
	_, err := c.Do("del", key)
	if err != nil {
		log.Println("del a key error, ", err)
		return false
	}
	return true
}

// Hash操作
func HashSet(key string, data map[string]interface{}) {
	defer c.Close()
	for k, v := range data {
		_, err := c.Do("hset", key, k, v)
		if err != nil {
			log.Println("hset a error, ", err)
			continue
		}
	}
}

func HashMGet(key string) map[string]interface{} {
	defer c.Close()

	data := make(map[string]interface{})

	reply, err := redis.ByteSlices(c.Do("hgetall", key))
	if err != nil {
		log.Println("hmget error, ", err)
		return nil
	}
	for i, v := range reply {
		fmt.Println(string(v))
		if i%2 == 0 {
			data[string(v)] = string(reply[i+1])
		}
		continue
	}
	return data

}
