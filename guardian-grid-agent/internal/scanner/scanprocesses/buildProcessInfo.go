package scanprocesses

import "github.com/shirou/gopsutil/v3/process"

func buildProcessInfo(p *process.Process) map[string]interface{} {
	name, _ := p.Name()
	if name == "" {
		return nil
	}

	pid := p.Pid
	cpu, _ := p.CPUPercent()
	memInfo, _ := p.MemoryInfo()
	username, _ := p.Username()

	rss := uint64(0)
	if memInfo != nil {
		rss = memInfo.RSS
	}

	return map[string]interface{}{
		"pid":         pid,
		"name":        name,
		"cpu_percent": cpu,
		"memory_rss":  rss,
		"user":        username,
	}
}
