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
	Labels   []Label
	SSH
}

type Label struct {
	key   string
	value string
}

type SSH struct {
	User     string
	Password string
}

const (
	defaultProtocol = "tcp"
	defaultTimeout  = 10 * time.Second
	defaultMaxConns = 1000
)

const (
	defaultUser     = "root"
	defaultPassword = "123456"
)

var defaultSSH = SSH{
	User:     defaultUser,
	Password: defaultPassword,
}

var serverSyncLabel = Label{
	key:   "sync",
	value: "",
}

var serverTypeLabel = Label{
	key:   "type",
	value: "worker",
}

func getDefaultLabels() []Label {
	var labels []Label

	labels = append(labels, serverTypeLabel, serverSyncLabel)

	return labels
}

type ServerOption func(*Server)

func Protocol(p string) ServerOption {
	return func(s *Server) {
		s.Protocol = p
	}
}
func Timeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.Timeout = timeout
	}
}
func MaxConns(maxconns int) ServerOption {
	return func(s *Server) {
		s.MaxConns = maxconns
	}
}
func TLS(tls *tls.Config) ServerOption {
	return func(s *Server) {
		s.TLS = tls
	}
}

func NewServer(addr string, port int, options ...func(*Server)) (*Server, error) {
	srv := Server{
		Addr:     addr,
		Port:     port,
		Protocol: defaultProtocol,
		Timeout:  defaultTimeout,
		MaxConns: defaultMaxConns,
		TLS:      nil,
		Labels:   getDefaultLabels(),
		SSH:      defaultSSH,
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
