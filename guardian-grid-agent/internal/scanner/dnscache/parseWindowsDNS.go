package dnscache

import "strings"

func parseWindowsDNS(raw string) []string {
	lines := strings.Split(raw, "\n")
	seen := make(map[string]bool)
	result := []string{}

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.Contains(line, "Record Name") {
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
