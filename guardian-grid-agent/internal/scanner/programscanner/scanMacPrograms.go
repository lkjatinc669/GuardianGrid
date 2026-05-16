//go:build darwin

package programscanner

import "strings"

func scanMacPrograms() (map[string]interface{}, error) {
	programs := []map[string]interface{}{}

	// 1. Applications folder (no reliable version)
	if data, err := runCmd("ls", "/Applications"); err == nil {
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			name := strings.TrimSpace(strings.TrimSuffix(line, ".app"))
			if name != "" && !containsNoise(name) {
				programs = append(programs, map[string]interface{}{
					"name": name,
				})
			}
		}
	}

	// 2. Homebrew (with version)
	if data, err := runCmd("brew", "list", "--versions"); err == nil {
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				programs = append(programs, map[string]interface{}{
					"name":    parts[0],
					"version": parts[1],
				})
			}
		}
	}

	return finalize(programs), nil
}
