package redis

import (
	"github.com/go-redis/redis"
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

func (redisServer *RedisServer) Subscribe() *redis.PubSub {
	return redisServer.client.Subscribe(redisServer.channelName)
}

func (redisServer *RedisServer) Publish(message string) error {
	return redisServer.client.Publish(redisServer.channelName, message).Err()
}
