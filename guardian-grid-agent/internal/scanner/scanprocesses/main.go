package scanprocesses

import (
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

// PUBLIC
func ScanProcesses() (map[string]interface{}, error) {
	procs, err := process.Processes()
	if err != nil {
		return emptyProcesses(), nil
	}

	result := make([]map[string]interface{}, 0)
	maxProcesses := 300 // safety limit

	for i, p := range procs {
		if i >= maxProcesses {
			break
		}

		info := buildProcessInfo(p)
		if info != nil {
			result = append(result, info)
		}
	}

	return map[string]interface{}{
		"processes": result,
		"count":     len(result),
		"timestamp": time.Now().Unix(),
	}, nil
}
