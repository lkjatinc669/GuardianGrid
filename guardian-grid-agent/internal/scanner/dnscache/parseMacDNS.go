//go:build darwin

package dnscache

import "strings"

func parseMacDNS(raw string) []string {
	lines := strings.Split(raw, "\n")
	seen := make(map[string]bool)
	result := []string{}

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "name:") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				domain := strings.TrimSpace(parts[1])

				if domain != "" && !seen[domain] {
					seen[domain] = true
					result = append(result, domain)
				}
			}
		}
	}

	return result
}
