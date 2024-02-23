package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

// attaching write function to logWriter struct, so logWritter will be implementation of Writer interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}

func CustomHTTPWritter() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("response is", resp)
	// response body is Body io.ReadCloser
	// so we can not directly print the response
	// fmt.Println("response.body is", resp.Body)

	bs := make([]byte, 99999)
	n, err := resp.Body.Read(bs)
	fmt.Println("got", n, "bytes", " error is", err)
	fmt.Print(string(bs))

	fmt.Println(io.Copy(os.Stdout, resp.Body))

	fmt.Println("With logwritter")
	lw := logWriter{}
	fmt.Println(io.Copy(lw, resp.Body))

}
