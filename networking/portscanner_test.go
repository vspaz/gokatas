package networking

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func testSetup(server *Server, t *testing.T) {
	conn, err := net.Dial(
		server.Protocol,
		fmt.Sprintf("%s:%d", server.Address, server.Port),
	)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	defer conn.Close()
}

func TestScanUDPOk(t *testing.T) {
	serverInfo := &Server{
		Protocol: "udp",
		Address:  "localhost",
		Port:     80,
	}
	go testSetup(serverInfo, t)
	_, err := Scan(serverInfo)
	assert.Equal(t, nil, err)
}

func TestScanTCPOK(t *testing.T) {
	serverInfo := &Server{
		Protocol: "tcp",
		Address:  "localhost",
		Port:     80,
	}
	go testSetup(serverInfo, t)
	_, err := Scan(serverInfo)
	assert.Equal(t, nil, err)
}
