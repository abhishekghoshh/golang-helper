package main

import (
	"log"
	"net/http"

	"github.com/abhishekghoshh/group-chat/pkg/config"
	"github.com/abhishekghoshh/group-chat/pkg/redis"
	"github.com/abhishekghoshh/group-chat/pkg/server"
)

func main() {
	cfg := config.New()

	serverAddr := cfg.GetString("server.address") + ":" + cfg.GetString("server.port")
	serverName := cfg.GetString("server.name")

	redisServer := redis.New(
		cfg.GetString("redis.host"),
		cfg.GetString("redis.port"),
		cfg.GetString("redis.channel"),
	)
	chatServer := server.New(serverName, redisServer)

	http.HandleFunc("/chat", chatServer.Chat)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
