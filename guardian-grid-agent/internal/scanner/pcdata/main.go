package pcdata

import (
	"runtime"
)

// PUBLIC
func ScanPCData() (map[string]interface{}, error) {
	h := getHostInfo()
	c := getCPUInfo()
	m := getMemInfo()

	return map[string]interface{}{
		"pc_data": map[string]interface{}{
			"hostname":       h.hostname,
			"os":             h.os,
			"platform":       h.platform,
			"platform_ver":   h.platformVer,
			"kernel_version": h.kernel,
			"architecture":   runtime.GOARCH,

			"cpu_model": c.model,
			"cpu_cores": c.cores,

			"ram_total_mb": m.totalMB,

			"boot_time": h.bootTime, // unix seconds
		},
	}, nil
}
