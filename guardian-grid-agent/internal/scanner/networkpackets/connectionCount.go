package networkpackets

import "github.com/shirou/gopsutil/v3/net"

func getConnCount() int {
	conns, err := net.Connections("inet")
	if err != nil {
		return 0
	}
	return len(conns)
}
