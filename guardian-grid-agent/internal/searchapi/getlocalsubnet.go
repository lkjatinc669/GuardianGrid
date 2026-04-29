package searchapi

import (
	"net"
)

func GetLocalSubnet() ([]string, error) {
	var ips []string

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range interfaces {
		addrs, _ := iface.Addrs()

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok || ipNet.IP.IsLoopback() {
				continue
			}

			ip := ipNet.IP.To4()
			if ip == nil {
				continue
			}

			baseIP := ip.Mask(ipNet.Mask)

			// generate /24 range (simple assumption)
			for i := 1; i < 255; i++ {
				newIP := make(net.IP, len(baseIP))
				copy(newIP, baseIP)
				newIP[3] = byte(i)
				ips = append(ips, newIP.String())
			}
		}
	}

	return ips, nil
}
