//go:build !windows && !linux && !darwin

package dnscache

func scanDNS() (map[string]interface{}, error) {
	return emptyDNS(), nil
}
