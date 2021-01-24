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
