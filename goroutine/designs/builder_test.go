package designs

import (
	"fmt"
	"testing"
	"time"
)

type Server struct {
	IP       string
	Port     int
	Timeout  time.Duration
	MaxConns int
}

type ServerBuilder struct {
	Server
}

func (sb *ServerBuilder) Create(ip string, port int) *ServerBuilder {
	sb.Server.IP = ip
	sb.Server.Port = port
	// for default value
	sb.Server.MaxConns = 6
	sb.Server.Timeout = 10 * time.Second
	return sb
}

func (sb *ServerBuilder) WithTimeout(timeout time.Duration) *ServerBuilder {
	sb.Server.Timeout = timeout
	return sb
}

func (sb *ServerBuilder) WithMaxConn(maxConn int) *ServerBuilder {
	sb.Server.MaxConns = maxConn
	return sb
}

func (sb *ServerBuilder) Build() Server {
	return sb.Server
}

func TestBuilder(t *testing.T) {
	sb := ServerBuilder{}
	server := sb.Create("127.0.0.1", 8090).
		WithTimeout(5 * time.Second).
		WithMaxConn(100).
		Build()
	fmt.Println(server.IP, server.Port, server.Timeout, server.MaxConns)
}

type Option func(*Server)

func TimeoutOption(timout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timout
	}
}

func MaxConnOption(maxConns int) Option {
	return func(s *Server) {
		s.MaxConns = maxConns
	}
}

func TestBuilderWithFunctional(t *testing.T) {
	server := Server{
		IP:       "127.0.0.1",
		Port:     8090,
		Timeout:  10 * time.Second,
		MaxConns: 10,
	}
	// 可以用pipline封装，这里不再重复
	TimeoutOption(6 * time.Second)(&server)
	MaxConnOption(200)(&server)
	fmt.Println(server.IP, server.Port, server.Timeout, server.MaxConns)
}
