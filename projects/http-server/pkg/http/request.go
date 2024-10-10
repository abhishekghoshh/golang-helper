package http

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"strings"
)

func processRequest(buff []byte) (*HTTPRequest, error) {
	reader := bufio.NewReader(bytes.NewBuffer(buff))

	startLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	startLine = strings.TrimSpace(startLine)
	request, err := parseRequestLine(startLine)
	if err != nil {
		return nil, err
	}
	// preparing the headers from the reader
	headers := map[string]string{}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		line = strings.TrimSpace(line)
		if line == "" {
			break
		}

		headerParts := strings.Split(line, ":")
		if len(headerParts) != 2 {
			continue
		}
		key, value := strings.TrimSpace(headerParts[0]), strings.TrimSpace(headerParts[1])
		headers[key] = value
	}
	request.Headers = headers

	// checking the content length header
	contentLength := 0
	if cv, ok := headers["Content-Length"]; ok {
		contentLength, err = strconv.Atoi(cv)
		if err != nil {
			return nil, err
		}
	}

	// preparing the body from the reader
	body := make([]byte, contentLength)
	if contentLength > 0 {
		if _, err := io.ReadFull(reader, body); err != nil {
			return nil, err
		}
	}
	request.Body = body

	return &request, nil
}

func parseRequestLine(s string) (HTTPRequest, error) {
	r := bufio.NewReader(bytes.NewBuffer([]byte(s)))

	method, err := r.ReadString(' ')
	if err != nil {
		return HTTPRequest{}, err
	}

	m, err := HTTPMethodFromString(strings.TrimSpace(method))
	if err != nil {
		return HTTPRequest{}, err
	}

	requestPath, err := r.ReadString(' ')
	if err != nil {
		return HTTPRequest{}, err
	}
	requestPath = strings.TrimSpace(requestPath)
	path, queryParams := refineRequestPath(requestPath)

	return HTTPRequest{
		Method:      m,
		Url:         requestPath,
		Version:     "HTTP/1.1",
		Path:        path,
		QueryParams: queryParams,
	}, nil
}

func refineRequestPath(requestPath string) (string, map[string][]string) {
	path, queryParams := requestPath, map[string][]string{}

	if strings.Contains(requestPath, "?") {
		pathParts := strings.Split(requestPath, "?")
		path = pathParts[0]

		if len(pathParts) > 1 {
			queryParts := strings.Split(pathParts[1], "&")
			for _, query := range queryParts {
				qPart := strings.Split(query, "=")

				if len(qPart) != 2 {
					continue
				}

				key, value := strings.TrimSpace(qPart[0]), strings.TrimSpace(qPart[1])

				queryValueList := []string{}
				if l, ok := queryParams[key]; ok {
					queryValueList = l
				}

				queryValueList = append(queryValueList, value)
				queryParams[key] = queryValueList
			}
		}
	}
	return path, queryParams
}
