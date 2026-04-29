package pcdata

import "github.com/shirou/gopsutil/v3/mem"

type memInfo struct {
	totalMB uint64
}

func getMemInfo() memInfo {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return memInfo{}
	}

	return memInfo{
		totalMB: vm.Total / 1024 / 1024,
	}
}
