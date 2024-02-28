package main

import (
	"di/server"
	"di/test"
	"testing"
)

func TestServer(t *testing.T) {
	logger := &test.MockLogger{}
	server := &server.Server{Logger: logger}
	//mock request to send in server
	if len(logger.Logs) != 1 {
		t.Errorf("expected 1 log message, got %d", len(logger.Logs))
	}
}
