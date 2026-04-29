package pcdata

import "github.com/shirou/gopsutil/v3/host"

type hostInfo struct {
	hostname    string
	os          string
	platform    string
	platformVer string
	kernel      string
	bootTime    int64
}

func getHostInfo() hostInfo {
	info, err := host.Info()
	if err != nil {
		return hostInfo{}
	}

	return hostInfo{
		hostname:    info.Hostname,
		os:          info.OS,
		platform:    info.Platform,
		platformVer: info.PlatformVersion,
		kernel:      info.KernelVersion,
		bootTime:    int64(info.BootTime),
	}
}
