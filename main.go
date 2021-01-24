package main

import (
	"fmt"
	"gokatas/networking"
)

func main() {
	serverInfo := &networking.Server{
		Protocol: "udp",
		Address:  "scanme.nmap.org",
		Port:     80,
	}
	_, err := networking.Scan(serverInfo)
	fmt.Println(err)
}
