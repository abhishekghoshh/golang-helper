package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisServer struct {
	channelName string
	client      *redis.Client
}

func New(server, port, channelName string) *RedisServer {
	redisAddress := server + ":" + port
	var redisClient = redis.NewClient(
		&redis.Options{
			Addr: redisAddress,
		})
	return &RedisServer{
		channelName: channelName,
		client:      redisClient,
	}
}

// SET stores a string value.
func (redisServer *RedisServer) Set(ctx context.Context, key string, value interface{}) error {
	return redisServer.client.Set(ctx, key, value, 0).Err()
}

// SET retrieve a string value.
func (redisServer *RedisServer) Get(ctx context.Context, key string) (string, error) {
	return redisServer.client.Get(ctx, key).Result()
}

// You can also easily store and retrieve a hash:
func (redisServer *RedisServer) HSet(ctx context.Context, key string, value []string) (int64, error) {
	return redisServer.client.HSet(ctx, key, value).Result()
}

func (redisServer *RedisServer) HGet(ctx context.Context, key, field string) (string, error) {
	return redisServer.client.HGet(ctx, key, field).Result()
}

func (redisServer *RedisServer) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return redisServer.client.HGetAll(ctx, key).Result()
}
func (redisServer *RedisServer) HGetAllAndParse(ctx context.Context, key string, structType interface{}) error {
	return redisServer.client.HGetAll(ctx, key).Scan(structType)
}

func (redisServer *RedisServer) SetCounter(ctx context.Context, key string) (string, error) {
	return redisServer.client.Set(ctx, key, "0", 0).Result()
}
func (redisServer *RedisServer) IncrCounter(ctx context.Context, key string) (int64, error) {
	return redisServer.client.Incr(ctx, key).Result()
}
func (redisServer *RedisServer) IncrCounterBy(ctx context.Context, key string, by int64) (int64, error) {
	return redisServer.client.IncrBy(ctx, key, by).Result()
}

func (redisServer *RedisServer) Subscribe(ctx context.Context) *redis.PubSub {
	return redisServer.client.Subscribe(ctx, redisServer.channelName)
}

func (redisServer *RedisServer) Publish(ctx context.Context, message string) error {
	return redisServer.client.Publish(ctx, redisServer.channelName, message).Err()
}
