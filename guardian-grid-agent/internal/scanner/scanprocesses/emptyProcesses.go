package scanprocesses

import "time"

func emptyProcesses() map[string]interface{} {
	return map[string]interface{}{
		"processes": []interface{}{},
		"count":     0,
		"timestamp": time.Now().Unix(),
	}
}
