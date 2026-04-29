package searchapi

import (
	"fmt"
	"net"
	"time"
)

func IsPortOpen(ip string, port int) bool {
	address := net.JoinHostPort(ip, fmt.Sprintf("%d", port))
	conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
