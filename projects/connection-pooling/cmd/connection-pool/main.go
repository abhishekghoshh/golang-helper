package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql" // import the MySQL driver
)

type ConnectionPool struct {
	db             *sql.DB     // the underlying connection pool
	maxConnections int         // the maximum number of connections in the pool
	numConnections int         // the current number of connections in the pool
	mutex          *sync.Mutex // the mutex to synchronize access to the connection pool
}

func NewConnectionPool(dsn string, maxConnections int) (*ConnectionPool, error) {
	// create a new connection pool
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// set the maximum number of connections in the pool
	db.SetMaxOpenConns(maxConnections)

	// create a new ConnectionPool instance
	p := &ConnectionPool{
		db:             db,
		maxConnections: maxConnections,
		numConnections: 0,
		mutex:          &sync.Mutex{},
	}
	return p, nil
}

// GetConnection acquires a connection from the pool
func (p *ConnectionPool) GetConnection() (*sql.DB, func(), error) {
	// acquire the mutex lock
	p.mutex.Lock()
	defer p.mutex.Unlock()
	// check if the pool is full
	if p.numConnections == p.maxConnections {
		return nil, func() {}, fmt.Errorf("connection pool is full")
	}

	// increment the number of connections in the pool
	p.numConnections++

	// ReleaseConnection releases a connection back to the pool
	release := func() {
		// acquire the mutex lock
		m := p.mutex
		m.Lock()
		defer m.Unlock()
		// decrement the number of connections in the pool
		p.numConnections--
	}
	// return a connection from the underlying pool
	return p.db, release, nil
}

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

func ping(pool *ConnectionPool) http.HandlerFunc {
	pingFunc := func(w http.ResponseWriter, r *http.Request) {
		// acquire a connection from the pool
		_, release, err := pool.GetConnection()
		defer release()
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			message := FailedMessage()
			json.NewEncoder(w).Encode(message)
			return
		}
		time.Sleep(1 * time.Second)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		message := SuccessfulMessage()
		json.NewEncoder(w).Encode(message)
	}
	return http.HandlerFunc(pingFunc)

}
func main() {
	// create a new connection pool
	pool, err := NewConnectionPool("abhishek:abhishek@tcp(localhost:3306)/user", 6)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.db.Close()
	// assigning the handler
	http.HandleFunc("/ping", ping(pool))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("There was an error listening on port :8080", err)
	}
}
