package liveactivity

import "github.com/shirou/gopsutil/v3/load"

type loadInfo struct {
	load1  float64
	load5  float64
	load15 float64
}

func getLoadAvg() loadInfo {
	avg, err := load.Avg()
	if err != nil {
		return loadInfo{}
	}

	return loadInfo{
		load1:  avg.Load1,
		load5:  avg.Load5,
		load15: avg.Load15,
	}
}
