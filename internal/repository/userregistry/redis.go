package userregistry

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisRegistry struct {
	client *redis.Client
}

func NewRedisRepository(redisURL string) *RedisRegistry {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "",
		DB:       0,
	})
	return &RedisRegistry{client: redisClient}
}

func (r *RedisRegistry) GetUserRegistry(userId string) (string, error) {
	value, err := r.client.Get(context.Background(), userId).Result()
	return value, err
}

func (r *RedisRegistry) RegisterUser(userId, registry string) error {
	// valueBytes, _ := json.Marshal(registry)
	return r.client.Set(context.Background(), userId, registry, 0).Err()
}

func (r *RedisRegistry) UnregisterUser(userId string) error {
	return r.client.Del(context.Background(), userId).Err()
}

func (r *RedisRegistry) Close() error {
	return r.client.Close()
}
