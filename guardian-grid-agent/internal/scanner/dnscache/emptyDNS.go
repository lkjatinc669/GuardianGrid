package dnscache

func emptyDNS() map[string]interface{} {
	return map[string]interface{}{
		"dns_cache": []string{},
		"count":     0,
	}
}
