package programscanner

func addProgram(item interface{}, result *[]map[string]interface{}, seen map[string]bool) {
	obj, ok := item.(map[string]interface{})
	if !ok {
		return
	}

	name, _ := obj["DisplayName"].(string)
	version, _ := obj["DisplayVersion"].(string)
	publisher, _ := obj["Publisher"].(string)

	if name == "" || containsNoise(name) {
		return
	}

	key := name + "|" + version
	if seen[key] {
		return
	}
	seen[key] = true

	entry := map[string]interface{}{
		"name":      name,
		"version":   version,
		"publisher": publisher,
	}

	*result = append(*result, entry)
}
