package networking

import (
	"fmt"
	"net"
)

type Server struct {
	protocol string
	host     string
	port     int
}

func Scan(server *Server) (net.Conn, error) {
	conn, err := net.Dial(
		server.protocol,
		fmt.Sprintf("%s:%d", server.host, server.port),
	)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}
	return conn, nil
}
