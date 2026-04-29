package networkpackets

import "github.com/shirou/gopsutil/v3/net"

type netDiff struct {
	bytesSent   uint64
	bytesRecv   uint64
	packetsSent uint64
	packetsRecv uint64
}

func calcNetDiff(start, end net.IOCountersStat) netDiff {
	return netDiff{
		bytesSent:   end.BytesSent - start.BytesSent,
		bytesRecv:   end.BytesRecv - start.BytesRecv,
		packetsSent: end.PacketsSent - start.PacketsSent,
		packetsRecv: end.PacketsRecv - start.PacketsRecv,
	}
}
