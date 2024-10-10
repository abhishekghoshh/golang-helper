package main

import (
	"flag"

	"github.com/abhishekghoshh/feather/pkg/http"
)

var (
	host = flag.String("h", "0.0.0.0", "sets the host")
	port = flag.Int("p", 8080, "sets the port")
)

func main() {
	http.NewTcpServerWithGracefullShutdown(*host, *port)
}
