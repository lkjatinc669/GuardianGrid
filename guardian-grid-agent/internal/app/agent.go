package app

import (
	"fmt"
	"sync"
	"time"

	"guardian-grid-agent/internal/auth"
	"guardian-grid-agent/internal/scanner/activeusers"
	"guardian-grid-agent/internal/scanner/dnscache"
	"guardian-grid-agent/internal/scanner/liveactivity"
	"guardian-grid-agent/internal/scanner/networkpackets"
	"guardian-grid-agent/internal/scanner/openports"
	"guardian-grid-agent/internal/scanner/pcdata"
	"guardian-grid-agent/internal/scanner/persistence"
	"guardian-grid-agent/internal/scanner/programscanner"
	"guardian-grid-agent/internal/scanner/scanprocesses"
	"guardian-grid-agent/internal/scanner/systemuptime"

	"guardian-grid-agent/internal/storage"
	"guardian-grid-agent/internal/utils"
)

func RunAgent(baseURL string) {
	// 1. Auth / Registration
	config, err := auth.GetConfig()
	if err != nil {
		fmt.Println("🔑 Registering agent with API...")
		config, err = auth.Register(baseURL)
		if err != nil {
			fmt.Printf("❌ Registration failed: %v. Retrying in 10s...\n", err)
			time.Sleep(10 * time.Second)
			RunAgent(baseURL)
			return
		}
		fmt.Printf("✅ Registered! ID: %s\n", config.AgentID)
	}

	buffer := storage.NewBuffer(200)

	// 🔁 scan loop
	go func() {
		for {
			runAllScanners(buffer)
			time.Sleep(1 * time.Second)
		}
	}()

	// ⏱ send loop
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		payload := buffer.FlushGrouped()

		if len(payload) == 0 {
			continue
		}

		url := baseURL + "/agent/data"
		go utils.SendData(url, payload, config.Token)
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
		"persistence":  persistence.ScanPersistence,
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

			buffer.AddGrouped(k, data)
		}(key, fn)
	}

	wg.Wait()
}
