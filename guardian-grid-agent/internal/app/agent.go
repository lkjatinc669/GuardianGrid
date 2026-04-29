package app

import (
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

func RunAgent(apiURL string) {
	buffer := storage.NewBuffer(200)

	// 🔁 continuous scanning loop (non-blocking)
	go func() {
		for {
			data := runAllScanners()

			if len(data) > 0 {
				buffer.Add(data)
			}

			time.Sleep(3 * time.Second) // scan interval
		}
	}()

	// ⏱ send every 15 seconds
	ticker := time.NewTicker(15 * time.Second)

	for range ticker.C {
		payload := buffer.Flush()
		if len(payload) == 0 {
			continue
		}

		go utils.SendData(apiURL, payload)
	}
}

func runAllScanners() map[string]interface{} {
	result := make(map[string]interface{})

	type scanResult struct {
		key  string
		data map[string]interface{}
	}

	ch := make(chan scanResult)

	// 🔥 launch all scanners concurrently
	go runScanner("active_users", activeusers.ScanActiveUsers, ch)
	go runScanner("dns_cache", dnscache.ScanDNSCache, ch)
	go runScanner("live_activity", liveactivity.ScanLiveActivity, ch)
	go runScanner("network", networkpackets.ScanNetworkPackets, ch)
	go runScanner("open_ports", openports.ScanOpenPorts, ch)
	go runScanner("pc_data", pcdata.ScanPCData, ch)
	go runScanner("programs", programscanner.ScanPrograms, ch)
	go runScanner("processes", scanprocesses.ScanProcesses, ch)
	go runScanner("uptime", systemuptime.ScanSystemUptime, ch)

	// ⚠️ optional (heavy)
	// go runScanner("folder", func() (map[string]interface{}, error) {
	//     return scanfolder.Scan("C:\\Users")
	// }, ch)

	// collect results
	for i := 0; i < 9; i++ {
		res := <-ch
		if res.data != nil {
			result[res.key] = res.data
		}
	}

	return result
}

func runScanner(
	name string,
	fn func() (map[string]interface{}, error),
	ch chan<- struct {
		key  string
		data map[string]interface{}
	},
) {
	data, err := fn()
	if err != nil {
		ch <- struct {
			key  string
			data map[string]interface{}
		}{name, nil}
		return
	}

	ch <- struct {
		key  string
		data map[string]interface{}
	}{name, data}
}
