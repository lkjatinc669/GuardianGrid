//go:build windows

package dnscache

import "strings"

func parseWindowsDNS(raw string) []map[string]interface{} {
	lines := strings.Split(raw, "\n")
	seen := make(map[string]bool)
	result := []map[string]interface{}{}

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.Contains(line, "Record Name") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				domain := strings.TrimSpace(parts[1])

				if domain != "" && !seen[domain] {
					seen[domain] = true
					result = append(result, map[string]interface{}{
						"hostname": domain,
						"ip":       "---", // ipconfig /displaydns doesn't easily provide IP in same line
						"type":     "A",
					})
				}
			}
		}
	}

	return result
}
