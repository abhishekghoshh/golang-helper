package test

type MockLogger struct {
	Logs []string
}

func (l *MockLogger) Log(message string) error {
	l.Logs = append(l.Logs, message)
	return nil
}
