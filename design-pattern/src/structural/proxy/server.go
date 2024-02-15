package proxy

type Server interface {
	handleRequest(string, string) (int, string)
}
