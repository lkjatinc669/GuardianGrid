package programscanner

import "encoding/json"

func parseProgramsJSON(data []byte) (map[string]interface{}, error) {
	var raw interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return emptyPrograms(), nil
	}

	programs := []map[string]interface{}{}
	seen := make(map[string]bool)

	switch v := raw.(type) {
	case []interface{}:
		for _, item := range v {
			addProgram(item, &programs, seen)
		}
	case map[string]interface{}:
		addProgram(v, &programs, seen)
	}

	return finalize(programs), nil
}
