package openports

import (
	"fmt"
	"net"
	"time"
)

func isPortOpen(host string, port int) bool {
	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))

	conn, err := net.DialTimeout("tcp", address, 300*time.Millisecond)
	if err != nil {
		return false
	}
	conn.Close()

	return true
}
