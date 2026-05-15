package networkpackets

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/net"
)

type ConnectionInfo struct {
	LocalAddr  string `json:"local_addr"`
	RemoteAddr string `json:"remote_addr"`
	Status     string `json:"status"`
	PID        int32  `json:"pid"`
}

func getDetailedConnections() []ConnectionInfo {
	conns, err := net.Connections("all")
	if err != nil {
		return nil
	}

	result := make([]ConnectionInfo, 0)
	for _, conn := range conns {
		// Filter for active/meaningful connections if needed, 
		// but "all" is better for forensics.
		
		remote := ""
		if conn.Raddr.IP != "" {
			remote = fmt.Sprintf("%s:%d", conn.Raddr.IP, conn.Raddr.Port)
		}

		result = append(result, ConnectionInfo{
			LocalAddr:  fmt.Sprintf("%s:%d", conn.Laddr.IP, conn.Laddr.Port),
			RemoteAddr: remote,
			Status:     conn.Status,
			PID:        conn.Pid,
		})
	}
	return result
}
