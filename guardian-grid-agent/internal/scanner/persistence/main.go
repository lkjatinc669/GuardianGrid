package persistence

import (
	"runtime"
)

// PersistenceItem represents a discovered persistence mechanism
type PersistenceItem struct {
	Type     string `json:"type"`      // e.g., "Registry", "Systemd", "Cron"
	Name     string `json:"name"`      // Name of the entry
	Path     string `json:"path"`      // Path to the executable or script
	Location string `json:"location"`  // Specific location (e.g., registry key or file path)
}

// PUBLIC
func ScanPersistence() (map[string]interface{}, error) {
	var items []PersistenceItem
	var err error

	switch runtime.GOOS {
	case "windows":
		items, err = scanWindowsPersistence()
	case "linux":
		items, err = scanLinuxPersistence()
	case "darwin":
		items, err = scanMacPersistence()
	default:
		items = []PersistenceItem{}
	}

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"persistence": items,
		"count":       len(items),
	}, nil
}
