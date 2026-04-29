package openports

import "sync"

func scanAllPorts(host string) []int {
	var wg sync.WaitGroup
	var mu sync.Mutex

	openPorts := make([]int, 0)

	// ⚙️ tune this carefully
	sem := make(chan struct{}, 500)

	for port := 1; port <= 65535; port++ {
		wg.Add(1)

		go func(p int) {
			defer wg.Done()

			sem <- struct{}{}
			defer func() { <-sem }()

			if isPortOpen(host, p) {
				mu.Lock()
				openPorts = append(openPorts, p)
				mu.Unlock()
			}
		}(port)
	}

	wg.Wait()
	return openPorts
}
