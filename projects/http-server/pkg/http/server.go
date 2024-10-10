package http

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

/*
Imporvements:
1. Limit the number of threads for handling the request (Use a threadpool)
2. Connection timeout
3. TCP backlog queue (How many connection you want to keep in your backlog)
*/
func NewTcpServer(host string, port int) {
	listenOn := fmt.Sprintf("%s:%d", host, port)
	// creating a listener object
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Printf("Server Started on http://%s:%d", host, port)
	// an infinite loop to handle the multiple requests
	for {
		// Accept is a syscall it will block and wait for the incoming connection
		conn, err := listener.Accept()
		if err != nil && !errors.Is(err, net.ErrClosed) {
			log.Printf("error: listen-loop: %s\n", err.Error())
			continue
		}
		// the process is a blocking call so make all the requests concurrent,
		// we have to make it run in a separate thread
		go process(conn)
	}
}

func NewTcpServerWithGracefullShutdown(host string, port int) {
	listenOn := fmt.Sprintf("%s:%d", host, port)
	ctx, cancel := context.WithCancel(context.Background())

	var lc net.ListenConfig
	listener, err := lc.Listen(ctx, "tcp", listenOn)
	if err != nil {
		log.Fatalf("error: listen: %s\n", err.Error())
	}
	log.Printf("Server Started on http://%s:%d", host, port)
	// creating a channel to capture the os signal
	osSignalChannel := make(chan os.Signal, 1)
	signal.Notify(osSignalChannel, syscall.SIGINT, syscall.SIGTERM)

	// separtate channel to close the listener
	go func() {
		defer listener.Close()
		// this is a blocking call
		<-osSignalChannel
		log.Println("Calling the context cancel")
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			// termination signal received, exiting listener loop
			log.Println("closing the infinite loop of listener Accept()")
			return

		default:
			conn, err := listener.Accept()
			if err != nil && !errors.Is(err, net.ErrClosed) {
				log.Printf("error: listen-loop: %s\n", err.Error())
				continue
			}
			go process(conn)
		}
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	buff := make([]byte, 4096)
	// Read call is also blocking call, so the process will not move forward from here
	_, err := conn.Read(buff)
	if err != nil {
		log.Fatal(err)
	}
	// here we will process the incoming request and do the business logic
	req, err := processRequest(buff)
	if err == nil {
		resp := []byte("HTTP/1.1 500 Internal Server Error\r\n\r\nInternal Server error\r\n")
		// Write is also a blocking call and it takes a byte array
		conn.Write(resp)
		return
	}

	log.Println(req.Url)
	log.Println(req.Path)
	log.Println(req.QueryParams)
	log.Println(req.Headers)
	log.Println(string(req.Body))

	// now we have to return the response
	// An HTTP response is made up of three parts, each separated by a CRLF (\r\n)
	// 1. Status line, 2. Zero or more headers, each ending with a CRLF, 3. Optional response body.
	resp := []byte("HTTP/1.1 200 OK\r\n\r\nHello world\r\n")
	// Write is also a blocking call and it takes a byte array
	conn.Write(resp)
}
