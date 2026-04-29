package app

import (
	"sync"
	"time"

	"guardian-grid-agent/internal/scanner/activeusers"
	"guardian-grid-agent/internal/scanner/dnscache"
	"guardian-grid-agent/internal/scanner/liveactivity"
	"guardian-grid-agent/internal/scanner/networkpackets"
	"guardian-grid-agent/internal/scanner/openports"
	"guardian-grid-agent/internal/scanner/pcdata"
	"guardian-grid-agent/internal/scanner/programscanner"
	"guardian-grid-agent/internal/scanner/scanprocesses"
	"guardian-grid-agent/internal/scanner/systemuptime"

	"guardian-grid-agent/internal/storage"
	"guardian-grid-agent/internal/utils"
)

func RunAgent(baseURL string) {
	buffer := storage.NewBuffer(200)

	// 🔁 scan loop
	go func() {
		for {
			runAllScanners(buffer)
			time.Sleep(3 * time.Second)
		}
	}()

	// ⏱ send loop
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		payload := buffer.FlushGrouped()

		for endpoint, data := range payload {
			if len(data) == 0 {
				continue
			}

			url := baseURL + "/" + endpoint
			go utils.SendData(url, data)
		}
	}
}

func runAllScanners(buffer *storage.Buffer) {
	var wg sync.WaitGroup

	scanners := map[string]func() (map[string]interface{}, error){
		"activeusers":  activeusers.ScanActiveUsers,
		"dnscache":     dnscache.ScanDNSCache,
		"liveactivity": liveactivity.ScanLiveActivity,
		"network":      networkpackets.ScanNetworkPackets,
		"openports":    openports.ScanOpenPorts,
		"pcdata":       pcdata.ScanPCData,
		"programs":     programscanner.ScanPrograms,
		"processes":    scanprocesses.ScanProcesses,
		"uptime":       systemuptime.ScanSystemUptime,
	}

	for key, fn := range scanners {
		wg.Add(1)

		go func(k string, f func() (map[string]interface{}, error)) {
			defer wg.Done()

			data, err := f()
			if err != nil || data == nil {
				return
			}

			buffer.AddGrouped(k, data) // 🔥 key change
		}(key, fn)
	}

	wg.Wait()
}
