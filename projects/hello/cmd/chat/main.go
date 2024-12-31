package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/abhishekghoshh/group-chat/pkg/config"
	"github.com/abhishekghoshh/group-chat/pkg/redis"
	"github.com/abhishekghoshh/group-chat/pkg/server"
	"github.com/abhishekghoshh/group-chat/pkg/utils"
)

func main() {
	cfg := config.New()

	serverAddr := cfg.GetString("server.address") + ":" + cfg.GetString("server.port")
	serverName := serverName(cfg)

	redisServer := redis.New(
		cfg.GetString("redis.host"),
		cfg.GetString("redis.port"),
		cfg.GetString("redis.channel"),
	)
	chatServer := server.New(serverName, redisServer)

	http.HandleFunc("/chat", chatServer.Chat)
	http.HandleFunc("/health", health)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
func health(w http.ResponseWriter, r *http.Request) {
	status := map[string]string{
		"health": "ok",
	}
	json.NewEncoder(w).Encode(status)
}

func serverName(cfg *config.Config) string {
	autoServerName := cfg.GetBool("server.autoServerName")
	var serverName string
	if !autoServerName {
		serverName = cfg.GetString("server.name")
	} else {
		podName := utils.KubernetesPodName()
		if podName != "" {
			serverName = podName
		} else {
			serverName = utils.RandomUserName()
		}
	}
	log.Println("server name is ", serverName)
	return serverName
}
