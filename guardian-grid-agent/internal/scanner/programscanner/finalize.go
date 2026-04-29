package programscanner

func finalize(programs []map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"installed_programs": programs,
		"count":              len(programs),
	}
}
