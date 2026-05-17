package persistence

// PersistenceItem represents a discovered persistence mechanism
type PersistenceItem struct {
	Type     string `json:"type"`     // e.g., "Registry", "Systemd", "Cron"
	Name     string `json:"name"`     // Name of the entry
	Path     string `json:"path"`     // Path to the executable or script
	Location string `json:"location"` // Specific location (e.g., registry key or file path)
}

// PUBLIC
func ScanPersistence() (map[string]interface{}, error) {
	items, err := scanPersistence()

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"persistence": items,
		"count":       len(items),
	}, nil
}
