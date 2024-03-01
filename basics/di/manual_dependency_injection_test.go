package di

import (
	"testing"
)

type MockLogger struct {
	Logs []string
}

func (l *MockLogger) Log(message string) error {
	l.Logs = append(l.Logs, message)
	return nil
}

func Test_DependencyInjection(t *testing.T) {
	logger := &MockLogger{}
	_ = &Server{Logger: logger}
	if len(logger.Logs) != 1 {
		t.Errorf("expected 1 log message, got %d", len(logger.Logs))
	}
}
