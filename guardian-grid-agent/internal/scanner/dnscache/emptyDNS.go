package dnscache

func emptyDNS() map[string]interface{} {
	return map[string]interface{}{
		"dns_entries": []map[string]interface{}{},
		"count":       0,
	}
}
