package config

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type EnvConfigRedis struct {
	RedisHost     string `mapstructure:"redis_host"`
	RedisPort     string `mapstructure:"redis_port"`
	RedisPassword string `mapstructure:"redis_password"`
}

var (
	RedisConfig EnvConfigRedis
)

var (
	cachingRedis *redis.Client
)

func InitRedisClient() {
	redisOptions := &redis.Options{
		Addr:     RedisConfig.RedisHost + ":" + RedisConfig.RedisPort,
		Password: RedisConfig.RedisPassword,
		DB:       0,
	}
	// if !RedisConfig.RedisDisableTLS {
	// 	redisOptions.TLSConfig = &tls.Config{}
	// }

	cachingRedis = redis.NewClient(redisOptions)
	if cachingRedis == nil {
		log.Fatal("Can't connect to redis")
		return
	}
	status, err := cachingRedis.Ping().Result()
	if err != nil {
		log.Fatal("Can't connect to redis, ping failed:", err)
		return
	}
	log.Printf("Redis Ping Status: %s", status)
}

func WriteCache(id string, data interface{}) error {
	dataByte, _ := json.Marshal(data)
	err := cachingRedis.Set(id, string(dataByte), 60000000000).Err()
	if err != nil {
		log.Println("cachingRedis.Set", err)
		return err
	}

	return err
}

func ReadCache(id string) ([]byte, error) {
	val, err := cachingRedis.Get(id).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	} else if err != nil {
		log.Println("cachingRedis.Get", err)
		return nil, err
	}
	return val, nil
}

func SetWithExp(id string, data interface{}, exp time.Duration) error {
	dataByte, _ := json.Marshal(data)
	err := cachingRedis.Set(id, string(dataByte), exp*time.Second).Err()
	if err != nil {
		log.Println("cachingRedis.Set", err)
		return err
	}

	return err
}

func DeleteCache(id string) error {
	err := cachingRedis.Del(id).Err()
	if err != nil {
		log.Println("RedisData.Del", err)
		return err
	}
	return nil
}
