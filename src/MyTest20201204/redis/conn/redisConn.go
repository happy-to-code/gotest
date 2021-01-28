package conn

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

func main() {
	// isSet := SetEx("k", "abc", 20)
	// fmt.Println("isSet:", isSet)
	value := Get("k")
	fmt.Printf("value:%s\n", value)

	value1 := Get("k")
	fmt.Printf("value:%s\n", value1)
}

const (
	redis_passwd  = ""
	redis_host    = "10.1.3.150:6379"
	redis_port    = ""
	redis_address = "10.1.3.150:6379"
)

func getConn() redis.Conn {
	conn, err := redis.Dial("tcp", redis_address)
	if err != nil {
		log.Println("redis connect err", err)
		panic(err)
	}
	return conn
}

var c = getConn()

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

var (
	pool *redis.Pool
)

func init() {
	pool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   0, // 0表示没有限制
		IdleTimeout: 1 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redis_host)
		},
	}
}
