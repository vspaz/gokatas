package networking

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func testSetup(server *Server, t *testing.T) {
	conn, err := net.Dial(
		server.protocol,
		fmt.Sprintf("%s:%d", server.host, server.port),
	)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	defer conn.Close()
}

func TestScanUDPOk(t *testing.T) {
	serverInfo := &Server{
		protocol: "udp",
		host:     "localhost",
		port:     80,
	}
	go testSetup(serverInfo, t)
	_, err := Scan(serverInfo)
	assert.Equal(t, nil, err)
}

func TestScanTCPOK(t *testing.T) {
	serverInfo := &Server{
		protocol: "tcp",
		host:     "localhost",
		port:     80,
	}
	go testSetup(serverInfo, t)
	_, err := Scan(serverInfo)
	assert.Equal(t, nil, err)
}
