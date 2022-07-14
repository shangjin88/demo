package main

//reference： https://coolshell.cn/articles/21146.html

import (
	"crypto/tls"
	"fmt"
	"time"
)

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

const (
	defaultProtocol = "tcp"
	defaultTimeout  = 10 * time.Second
	defaultMaxConns = 1000
)

type option func(*Server)

func Protocol(p string) option {
	return func(s *Server) {
		s.Protocol = p
	}
}
func Timeout(timeout time.Duration) option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}
func MaxConns(maxconns int) option {
	return func(s *Server) {
		s.MaxConns = maxconns
	}
}
func TLS(tls *tls.Config) option {
	return func(s *Server) {
		s.TLS = tls
	}
}

func NewServer(addr string, port int, options ...func(server *Server)) (*Server, error) {
	srv := Server{
		Addr:     addr,
		Port:     port,
		Protocol: defaultProtocol,
		Timeout:  defaultTimeout,
		MaxConns: defaultMaxConns,
		TLS:      nil,
	}

	for _, option := range options {
		option(&srv)
	}

	return &srv, nil
}

func main() {
	srv, _ := NewServer("127.0.0.1", 8080, Protocol("udp"), Timeout(5*time.Second))
	fmt.Printf("srv: %v", srv)
}
