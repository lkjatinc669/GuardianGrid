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
		"dns_cache_raw": string(out),
		"count":         0,
	}, nil
}
