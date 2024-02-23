package http

import (
	"fmt"
	"net/http"
	"time"
)

// https://gobyexample.com/context
func Context() {
	http.HandleFunc("/hello", helloRoute)
	http.ListenAndServe(":8090", nil)
}

func helloRoute(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}
