package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Handles WebSocket connection
func handleTerminal(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return err
	}
	defer conn.Close()

	for {
		// Read command from WebSocket
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		cmdStr := string(message)
		fmt.Println("Executing:", cmdStr)

		// Create command execution instance
		cmd := exec.Command("docker", "exec", "-i", "my_container", "sh", "-c", cmdStr)

		// Get stdout and stderr pipes
		stdout, _ := cmd.StdoutPipe()
		stderr, _ := cmd.StderrPipe()
		if err := cmd.Start(); err != nil {
			log.Println("Failed to start command:", err)
			return err
		}

		// Stream stdout
		go func() {
			scanner := bufio.NewScanner(stdout)
			for scanner.Scan() {
				conn.WriteMessage(websocket.TextMessage, []byte(scanner.Text()+"\r\n"))
			}
		}()

		// Stream stderr
		go func() {
			scanner := bufio.NewScanner(stderr)
			for scanner.Scan() {
				conn.WriteMessage(websocket.TextMessage, []byte(scanner.Text()+"\r\n"))
			}
		}()

		// Wait for command to finish
		err = cmd.Wait()
		if err != nil {
			log.Println("Command execution error:", err)
		}
	}

	return nil
}

func main() {
	e := echo.New()

	e.Static("/", "static") // Serve `index.html`

	e.GET("/ws", handleTerminal)

	log.Println("Server running on http://localhost:8080")
	e.Logger.Fatal(e.Start(":8080"))
}
