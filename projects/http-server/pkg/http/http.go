package http

import (
	"fmt"
	"net/http"
)

type (
	HTTPMethod      int
	HTTPMessageType int
)

const (
	HTTPMessageRequest HTTPMessageType = iota
	HTTPMessageResponse
)

const (
	GET HTTPMethod = iota
	POST
	PUT
	PATCH
	DELETE
	CONNECT
	OPTIONS
)

type HTTPStatus struct {
	Status       int
	ReasonPhrase string
	Version      string
}

type HTTPRequest struct {
	Method      HTTPMethod
	Path        string
	Version     string
	Url         string
	QueryParams map[string][]string
	Status      HTTPStatus
	Headers     map[string]string
	Body        []byte
}

func NewHTTPStatus(status int) HTTPStatus {
	return HTTPStatus{
		Status:       status,
		ReasonPhrase: http.StatusText(status),
		Version:      "HTTP/1.1",
	}
}

func (h HTTPMethod) String() string {
	m := map[HTTPMethod]string{
		GET:     "GET",
		POST:    "POST",
		PUT:     "PUT",
		PATCH:   "PATCH",
		DELETE:  "DELETE",
		CONNECT: "CONNECT",
		OPTIONS: "OPTIONS",
	}

	return m[h]
}

func HTTPMethodFromString(s string) (HTTPMethod, error) {
	m := map[string]HTTPMethod{
		"GET":     GET,
		"POST":    POST,
		"PUT":     PUT,
		"PATCH":   PATCH,
		"DELETE":  DELETE,
		"CONNECT": CONNECT,
		"OPTIONS": OPTIONS,
	}

	method, ok := m[s]
	if !ok {
		return HTTPMethod(-1),
			fmt.Errorf("error: method: parse: invalid method string '%s'", s)
	}

	return method, nil
}
