package networkpackets

import (
	"time"

	"github.com/shirou/gopsutil/v3/net"
)

// PUBLIC
func ScanNetworkPackets() (map[string]interface{}, error) {
	start, err := net.IOCounters(false)
	if err != nil || len(start) == 0 {
		return emptyNetwork(), nil
	}

	// ultra-short window for high-frequency updates
	time.Sleep(500 * time.Millisecond)

	end, err := net.IOCounters(false)
	if err != nil || len(end) == 0 {
		return emptyNetwork(), nil
	}

	diff := calcNetDiff(start[0], end[0])
	conns := getDetailedConnections()

	return map[string]interface{}{
		"network_activity": map[string]interface{}{
			"bytes_sent":   diff.bytesSent,
			"bytes_recv":   diff.bytesRecv,
			"packets_sent": diff.packetsSent,
			"packets_recv": diff.packetsRecv,
			"connections":  conns,
			"count":        len(conns),
			"interval_sec": 2,
			"timestamp":    time.Now().Unix(),
		},
	}, nil
}
