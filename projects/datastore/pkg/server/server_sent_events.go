package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type SSEApi struct {
}

func NewSSEApi() Server {
	return &SSEApi{}
}

func (api *SSEApi) Init(e *echo.Echo) {
	e.GET("/http/sse", api.httpSSEConnection)
	e.GET("/http/sse/client-aware", api.clientAwareHttpSSEConnection)

}

// HTTP SSE API
func (api *SSEApi) httpSSEConnection(c echo.Context) error {
	return api.httpSSEHandler(c.Response(), c.Request())
}

func (api *SSEApi) httpSSEHandler(w http.ResponseWriter, r *http.Request) error {
	// Set CORS headers to allow all origins. You may want to restrict this to specific origins in a production environment.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Simulate sending events (you can replace this with real data)
	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "data: %s\n\n", fmt.Sprintf("Event %d", i))
		time.Sleep(500 * time.Millisecond)
		w.(http.Flusher).Flush()
	}
	return r.Context().Err()
}

// client cancellation aware sse connection
func (api *SSEApi) clientAwareHttpSSEConnection(c echo.Context) error {
	return api.clientAwarehttpSSEHandler(c.Response(), c.Request())
}

func (api *SSEApi) clientAwarehttpSSEHandler(w *echo.Response, r *http.Request) error {
	// Set http headers required for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// You may need this locally for CORS requests
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Create a channel for client disconnection
	clientGone := r.Context().Done()

	rc := http.NewResponseController(w)
	t := time.NewTicker(time.Second)
	defer t.Stop()
	for {
		select {
		case <-clientGone:
			err := r.Context().Err()
			fmt.Println("Client disconnected : " + err.Error())
			return err
		case <-t.C:
			// Send an event to the client
			// Here we send only the "data" field, but there are few others
			if _, err := fmt.Fprintf(w, "data: The time is %s\n\n", time.Now().Format(time.UnixDate)); err != nil {
				return err
			}
			if err := rc.Flush(); err != nil {
				return err
			}
		}
	}
}
