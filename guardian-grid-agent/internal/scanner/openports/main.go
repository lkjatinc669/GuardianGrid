package openports

import (
	"sort"
)

// PUBLIC
func ScanOpenPorts() (map[string]interface{}, error) {
	open := scanAllPorts("127.0.0.1")

	sort.Ints(open)

	return map[string]interface{}{
		"open_ports": open,
		"count":      len(open),
	}, nil
}
