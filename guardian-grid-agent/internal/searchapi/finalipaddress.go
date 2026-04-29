package searchapi

import (
	"fmt"

	"guardian-grid-agent/internal/utils"
)

func FinalIPAddress() {
	const PORT = 64289

	hosts := ScanForPort(PORT)

	if len(hosts) == 0 {
		fmt.Println("No devices found with port open")
		return
	}

	for _, ip := range hosts {
		url := fmt.Sprintf("http://%s:%d/gg", ip, PORT)

		// assuming GetData(url string) (string, error)
		resp, err := utils.GetData(url, map[string]interface{}{})
		if err != nil {
			fmt.Println("Error fetching from", ip, ":", err)
			continue
		}

		fmt.Println("Response from", ip, ":", resp)
	}
}
