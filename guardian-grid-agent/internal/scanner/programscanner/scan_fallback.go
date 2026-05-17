//go:build !windows && !linux && !darwin

package programscanner

func scanPrograms() (map[string]interface{}, error) {
	return emptyPrograms(), nil
}
