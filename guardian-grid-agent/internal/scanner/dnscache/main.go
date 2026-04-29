package dnscache

import (
	"runtime"
)

// PUBLIC
func ScanDNSCache() (map[string]interface{}, error) {
	switch runtime.GOOS {
	case "windows":
		return scanWindowsDNS()
	case "linux":
		return scanLinuxDNS()
	case "darwin":
		return scanMacDNS()
	default:
		return emptyDNS(), nil
	}
}
