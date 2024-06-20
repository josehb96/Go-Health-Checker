package main

import (
	"fmt"
	"net"
	"time"
)

func Check(destination, port string) string {

	address := destination + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout) // Establish a connection to the address and port specified above with a timeout of 5 seconds
	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable, \n Error: %v", destination, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable,\n From: %v\n To: %v", destination, conn.LocalAddr(), conn.RemoteAddr())
	}

	return status

}
