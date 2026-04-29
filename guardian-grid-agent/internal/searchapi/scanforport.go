package searchapi

import (
	"sync"
)

func ScanForPort(port int) []string {
	ips, _ := GetLocalSubnet()

	var results []string
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, ip := range ips {
		wg.Add(1)

		go func(ip string) {
			defer wg.Done()

			if IsPortOpen(ip, port) {
				mu.Lock()
				results = append(results, ip)
				mu.Unlock()
			}
		}(ip)
	}

	wg.Wait()
	return results
}
