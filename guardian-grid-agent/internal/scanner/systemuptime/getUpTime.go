package systemuptime

import (
	"time"

	"github.com/shirou/gopsutil/v3/host"
)

func getUptime() (uint64, int64) {
	info, err := host.Info()
	if err != nil {
		return 0, 0
	}

	now := time.Now().Unix()
	boot := int64(info.BootTime)

	uptime := uint64(now - boot)

	return uptime, boot
}
