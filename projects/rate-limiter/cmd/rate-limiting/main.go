package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/didip/tollbooth"
	"golang.org/x/time/rate"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func SuccessfulMessage() *Message {
	return &Message{
		Status: "Successful",
		Body:   "Hi! This is a succesfull message",
	}
}
func FailedMessage() *Message {
	return &Message{
		Status: "Request Failed",
		Body:   "The API is at capacity, try again later.",
	}
}

func ping(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	message := SuccessfulMessage()
	json.NewEncoder(writer).Encode(message)
}

func rateLimiter(next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	// NewLimiter(r, b) returns a new limiter that allows events up to rate r and permits bursts of at most b tokens.
	limiter := rate.NewLimiter(2, 4)
	// NewLimiter() returns a struct that will allow an average of r requests per second and bursts,
	// which are strings of consecutive requests, of up to b. Now that we understand that,
	// the rest follows the standard middleware pattern.
	// The middleware creates the limiter struct and then returns a handler created from an anonymous struct.

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if !limiter.Allow() {
				message := FailedMessage()
				w.WriteHeader(http.StatusTooManyRequests)
				json.NewEncoder(w).Encode(message)
			} else {
				next(w, r)
			}
		})
}

func perClientRateLimiter(next func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}
	var (
		mu      sync.Mutex
		clients = make(map[string]*client)
	)
	go func() {
		for {
			time.Sleep(time.Minute)
			// Lock the mutex to protect this section from race conditions.
			mu.Lock()
			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}
			mu.Unlock()
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the IP address from the request.
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Lock the mutex to protect this section from race conditions.
		mu.Lock()
		if _, found := clients[ip]; !found {
			clients[ip] = &client{limiter: rate.NewLimiter(2, 4)}
		}
		clients[ip].lastSeen = time.Now()
		if !clients[ip].limiter.Allow() {
			mu.Unlock()
			message := FailedMessage()
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(message)
			return
		}
		mu.Unlock()
		next(w, r)
	})
}

func main() {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/ping/rate-limit", rateLimiter(ping))
	http.HandleFunc("/ping/rate-limit-per-client", perClientRateLimiter(ping))

	jsonMessage, _ := json.Marshal(FailedMessage())
	// Create a request limiter per handler.
	tlbthLimiter := tollbooth.NewLimiter(5, nil)
	tlbthLimiter.SetMessageContentType("application/json")
	tlbthLimiter.SetMessage(string(jsonMessage))
	http.Handle("/ping/rate-limit-tollbooth", tollbooth.LimitFuncHandler(tlbthLimiter, ping))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("There was an error listening on port :8080", err)
	}
}
