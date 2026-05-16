//go:build linux

package dnscache

func scanLinuxDNS() (map[string]interface{}, error) {
	// systemd-resolved (most modern distros)
	out, err := runCmd("resolvectl", "query", "google.com")
	if err != nil {
		return emptyDNS(), nil
	}

	// Linux DNS cache is not directly accessible
	// so return minimal signal
	return map[string]interface{}{
		"dns_entries": []map[string]interface{}{
			{
				"hostname": "Raw Output",
				"ip":       "N/A",
				"type":     "RAW",
				"details":  string(out),
			},
		},
		"count": 1,
	}, nil
}
