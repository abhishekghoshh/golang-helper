package server

import (
	"log"
	"net/http"
	"time"

	"github.com/abhishekghoshh/group-chat/pkg/redis"
	"github.com/goombaio/namegenerator"
	"github.com/gorilla/websocket"
)

type ChatServer struct {
	serverName string
	redis      *redis.RedisServer
}

func New(serverName string, redis *redis.RedisServer) *ChatServer {
	return &ChatServer{
		serverName: serverName,
		redis:      redis,
	}
}

var upgrader = websocket.Upgrader{}

func (chatServer *ChatServer) Chat(w http.ResponseWriter, r *http.Request) {
	subscriber := chatServer.redis.Subscribe()

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	// genarating unique name for the user
	name := chatServer.randomUserName()

	welcomeMessage := []byte("Hi " + name + ", Welcome to server : " + chatServer.serverName)
	if err := c.WriteMessage(1, welcomeMessage); err != nil {
		return
	}
	notificationMessage := name + " has joined from " + chatServer.serverName
	if err := chatServer.redis.Publish(notificationMessage); err != nil {
		log.Println("error in publishing message", err)
		return
	}

	ch := make(chan bool)

	// incoming messge on a different thread
	go func() {
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("error in receiving message", err)
				ch <- true
				return
			}
			sendingMessage := name + " : " + string(msg)
			if err := chatServer.redis.Publish(sendingMessage); err != nil {
				log.Println("error in publishing message", err)
				ch <- true
				return
			}
		}
	}()

	// outgoing request on a different thread
	go func() {
		for {
			msg, err := subscriber.ReceiveMessage()
			if err != nil {
				log.Println("error in reciving message from redis", err)
				ch <- true
				return
			}
			if err = c.WriteMessage(1, []byte(msg.Payload)); err != nil {
				log.Println("error in returning response", err)
				ch <- true
				return
			}
		}
	}()

	// blocking call to hold the connection
	<-ch
}

func (*ChatServer) randomUserName() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	name := nameGenerator.Generate()
	return name
}
