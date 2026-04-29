package scanner

// import (
// 	"github.com/shirou/gopsutil/v3/process"

// 	"guardian-grid-agent/utils"
// )

// // ---- structure ----
// type ProcInfo struct {
// 	PID  int32  `json:"pid"`
// 	Name string `json:"name"`
// 	Path string `json:"path"`
// }

// // ---- main function ----
// func SendProcesses(batchSize int) {
// 	procs, _ := process.Processes()

// 	var batch []ProcInfo

// 	for _, p := range procs {
// 		name, _ := p.Name()
// 		exe, _ := p.Exe()

// 		info := ProcInfo{
// 			PID:  p.Pid,
// 			Name: name,
// 			Path: exe,
// 		}

// 		batch = append(batch, info)

// 		if len(batch) >= batchSize {
// 			utils.SendData("http://your-api/processes", batch)
// 			batch = nil
// 		}
// 	}

// 	// send remaining
// 	if len(batch) > 0 {
// 		utils.SendData("http://your-api/processes", batch)
// 	}
// }
