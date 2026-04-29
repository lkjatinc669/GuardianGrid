package networkpackets

import "time"

func emptyNetwork() map[string]interface{} {
	return map[string]interface{}{
		"network_activity": map[string]interface{}{
			"bytes_sent":   0,
			"bytes_recv":   0,
			"packets_sent": 0,
			"packets_recv": 0,
			"connections":  0,
			"interval_sec": 0,
			"timestamp":    time.Now().Unix(),
		},
	}
}
