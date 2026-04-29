package liveactivity

import (
	"runtime"
	"time"
)

// PUBLIC
func ScanLiveActivity() (map[string]interface{}, error) {
	cpuPercent := getCPUPercent()
	memStats := getMemoryStats()
	procCount := getProcessCount()
	loadAvg := getLoadAvg()

	return map[string]interface{}{
		"live_activity": map[string]interface{}{
			"cpu_percent":     cpuPercent,
			"memory_percent":  memStats.percent,
			"memory_used_mb":  memStats.usedMB,
			"memory_total_mb": memStats.totalMB,
			"process_count":   procCount,
			"goroutines":      runtime.NumGoroutine(),
			"load_1m":         loadAvg.load1,
			"load_5m":         loadAvg.load5,
			"load_15m":        loadAvg.load15,
			"timestamp":       time.Now().Unix(),
		},
	}, nil
}
