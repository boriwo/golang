//
// Golang Workshop 2024
//

package main

import (
	"fmt"
)

type Server struct {
	Host    string
	Port    int
	Timeout int
}

type ServerOption func(*Server)

func WithHost(host string) ServerOption {
	return func(s *Server) {
		s.Host = host
	}
}

func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.Port = port
	}
}

func WithTimeout(timeout int) ServerOption {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func NewServer(options ...ServerOption) *Server {
	server := &Server{
		Host:    "localhost",
		Port:    8080,
		Timeout: 30,
	}
	for _, option := range options {
		option(server)
	}
	return server
}

func main() {
	server := NewServer(
		WithHost("example.com"),
		WithPort(9090),
		WithTimeout(60),
	)
	fmt.Printf("Server Configuration: Host=%s, Port=%d, Timeout=%d\n", server.Host, server.Port, server.Timeout)
}
