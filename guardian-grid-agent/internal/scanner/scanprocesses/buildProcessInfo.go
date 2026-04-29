package scanprocesses

import "github.com/shirou/gopsutil/v3/process"

func buildProcessInfo(p *process.Process) map[string]interface{} {
	name, _ := p.Name()
	if name == "" {
		return nil
	}

	pid := p.Pid
	cpu, _ := p.CPUPercent()
	mem, _ := p.MemoryPercent()
	username, _ := p.Username()

	return map[string]interface{}{
		"pid":    pid,
		"name":   name,
		"cpu":    cpu,
		"memory": mem,
		"user":   username,
	}
}
