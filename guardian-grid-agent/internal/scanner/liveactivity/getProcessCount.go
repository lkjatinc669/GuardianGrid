package liveactivity

import "github.com/shirou/gopsutil/v3/process"

func getProcessCount() int {
	procs, err := process.Processes()
	if err != nil {
		return 0
	}
	return len(procs)
}
