package dnscache

// PUBLIC
func ScanDNSCache() (map[string]interface{}, error) {
	return scanDNS()
}
