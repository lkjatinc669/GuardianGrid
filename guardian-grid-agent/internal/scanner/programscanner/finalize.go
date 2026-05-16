package programscanner

func finalize(programs []map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"programs": programs,
		"count":    len(programs),
	}
}
