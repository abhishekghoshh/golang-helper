package server

import (
	"context"
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	subscriber := chatServer.redis.Subscribe(ctx)
	defer subscriber.Close()

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
	if err := chatServer.redis.Publish(ctx, notificationMessage); err != nil {
		log.Println("error in publishing message", err)
		return
	}

	// currently we have only read channel
	// so if there is any error in reading then it will close the channel as well as subscriber
	// so the http request and write thread will be closed, this is one way
	// we could also use a writing channel then writing thread will also able to close the reading thread
	readCh := make(chan bool)

	// incoming messge on a different thread
	go func() {
		defer func() {
			readCh <- true
		}()
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("closing the reading thread on error in receiving message", err)
				return
			}
			sendingMessage := name + " : " + string(msg)
			if err := chatServer.redis.Publish(ctx, sendingMessage); err != nil {
				log.Println("closing the reading thread on error in publishing message", err)
				return
			}
		}

	}()

	// outgoing request on a different thread
	go func() {
		for {
			msg, err := subscriber.ReceiveMessage(ctx)
			if err != nil {
				log.Println("closing the writing thread on error in reciving message from redis", err)
				return
			}
			if err = c.WriteMessage(1, []byte(msg.Payload)); err != nil {
				log.Println("closing the writing thread on error in returning response", err)
				return
			}
		}
	}()

	// blocking call to hold the connection
	<-readCh
}

func (*ChatServer) randomUserName() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	name := nameGenerator.Generate()
	return name
}
