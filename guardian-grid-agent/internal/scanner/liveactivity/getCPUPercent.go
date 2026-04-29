package liveactivity

import "github.com/shirou/gopsutil/v3/cpu"

func getCPUPercent() float64 {
	// non-blocking snapshot
	percent, err := cpu.Percent(0, false)
	if err != nil || len(percent) == 0 {
		return 0
	}
	return percent[0]
}
