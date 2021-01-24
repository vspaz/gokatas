package networking

import (
	"fmt"
	"net"
)

type Server struct {
	Protocol string
	Address  string
	Port     int
}

func Scan(server *Server) (net.Conn, error) {
	conn, err := net.Dial(
		server.Protocol,
		fmt.Sprintf("%s:%d", server.Address, server.Port),
	)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}
	return conn, nil
}
