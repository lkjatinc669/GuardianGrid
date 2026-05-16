package programscanner

import (
	"runtime"
)

// PUBLIC
func ScanPrograms() (map[string]interface{}, error) {
	switch runtime.GOOS {
	case "windows":
		return scanWindowsPrograms()
	case "linux":
		return scanLinuxPrograms()
	case "darwin":
		return scanMacPrograms()
	default:
		return map[string]interface{}{
			"programs": []interface{}{},
			"count":    0,
		}, nil
	}
}
