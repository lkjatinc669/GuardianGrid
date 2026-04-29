package programscanner

func emptyPrograms() map[string]interface{} {
	return map[string]interface{}{
		"installed_programs": []interface{}{},
		"count":              0,
	}
}
