package programscanner

import "strings"

func containsNoise(name string) bool {
	keywords := []string{
		"Debug", "Symbols", "Test Suite",
		"Development Libraries", "Documentation", "Bootstrap",
	}

	for _, k := range keywords {
		if strings.Contains(name, k) {
			return true
		}
	}
	return false
}
