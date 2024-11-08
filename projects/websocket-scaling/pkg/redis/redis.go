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
	var redisClient = redis.NewClient(&redis.Options{
		Addr: redisAddress,
	})
	return &RedisServer{
		client: redisClient,
	}
}

func (redisServer *RedisServer) Subscribe(ctx context.Context) *redis.PubSub {
	return redisServer.client.Subscribe(ctx, redisServer.channelName)
}

func (redisServer *RedisServer) Publish(ctx context.Context, message string) error {
	return redisServer.client.Publish(ctx, redisServer.channelName, message).Err()
}
