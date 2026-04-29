package systemuptime

import (
	"time"
)

// PUBLIC
func ScanSystemUptime() (map[string]interface{}, error) {
	uptimeSec, bootTime := getUptime()

	return map[string]interface{}{
		"system_uptime": map[string]interface{}{
			"uptime_seconds": uptimeSec,
			"uptime_hours":   uptimeSec / 3600,
			"boot_time":      bootTime, // unix timestamp
			"timestamp":      time.Now().Unix(),
		},
	}, nil
}
