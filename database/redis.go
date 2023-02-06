package database

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func RedisConnection() *redis.Client {
	if RedisClient == nil {
		redis := redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_HOST"),
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		RedisClient = redis
	}
	return RedisClient
}

func SetCache(key string, value any, ttl time.Duration) {
	RedisConnection()
	jsontr, _ := json.Marshal(value)
	err := RedisClient.Set(Ctx, key, string(jsontr), ttl*time.Minute).Err()
	log.Error().Str(key, "set data").Msg("0000")
	if err != nil {
		log.Error().Str(key, "this key value not set, error on set cache").Msg(err.Error())
	}
}

func GetCache(key string) interface{} {
	RedisConnection()
	val, err := RedisClient.Get(Ctx, key).Result()
	if err != nil {
		log.Error().Str(key, "data not found").Msg(err.Error())
		return nil
	}
	if val == "" {
		log.Error().Str(key, "data not found")
		return nil
	}
	var v interface{}
	json.Unmarshal([]byte(val), &v)
	return v
}
