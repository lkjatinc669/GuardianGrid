package pcdata

import "github.com/shirou/gopsutil/v3/cpu"

type cpuInfo struct {
	model string
	cores int32
}

func getCPUInfo() cpuInfo {
	info, err := cpu.Info()
	if err != nil || len(info) == 0 {
		return cpuInfo{}
	}

	return cpuInfo{
		model: info[0].ModelName,
		cores: info[0].Cores,
	}
}
