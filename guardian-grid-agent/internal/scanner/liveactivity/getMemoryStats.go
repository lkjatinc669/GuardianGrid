package liveactivity

import "github.com/shirou/gopsutil/v3/mem"

type memInfo struct {
	percent float64
	usedMB  float64
	totalMB float64
}

func getMemoryStats() memInfo {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return memInfo{}
	}

	return memInfo{
		percent: vm.UsedPercent,
		usedMB:  float64(vm.Used) / 1024 / 1024,
		totalMB: float64(vm.Total) / 1024 / 1024,
	}
}
