package dnscache

func scanWindowsDNS() (map[string]interface{}, error) {
	out, err := runCmd("ipconfig", "/displaydns")
	if err != nil {
		return emptyDNS(), nil
	}

	domains := parseWindowsDNS(string(out))

	return map[string]interface{}{
		"dns_entries": domains,
		"count":       len(domains),
	}, nil
}
