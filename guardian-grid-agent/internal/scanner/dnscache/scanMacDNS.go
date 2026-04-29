package dnscache

func scanMacDNS() (map[string]interface{}, error) {
	out, err := runCmd("dscacheutil", "-cachedump", "-entries", "Host")
	if err != nil {
		return emptyDNS(), nil
	}

	domains := parseMacDNS(string(out))

	return map[string]interface{}{
		"dns_cache": domains,
		"count":     len(domains),
	}, nil
}
