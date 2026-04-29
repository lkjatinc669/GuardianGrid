package programscanner

import "strings"

func parseKeyValueLines(data []byte) (map[string]interface{}, error) {
	lines := strings.Split(string(data), "\n")
	programs := []map[string]interface{}{}
	seen := make(map[string]bool)

	for _, line := range lines {
		parts := strings.Split(line, "\t")
		if len(parts) < 1 {
			continue
		}

		name := strings.TrimSpace(parts[0])
		version := ""
		if len(parts) > 1 {
			version = strings.TrimSpace(parts[1])
		}

		if name == "" || containsNoise(name) {
			continue
		}

		key := name + "|" + version
		if seen[key] {
			continue
		}
		seen[key] = true

		entry := map[string]interface{}{"name": name}
		if version != "" {
			entry["version"] = version
		}

		programs = append(programs, entry)
	}

	return finalize(programs), nil
}
