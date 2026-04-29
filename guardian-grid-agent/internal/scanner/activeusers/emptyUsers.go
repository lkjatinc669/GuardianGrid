package activeusers

import "time"

func emptyUsers() map[string]interface{} {
	return map[string]interface{}{
		"active_users": []interface{}{},
		"count":        0,
		"timestamp":    time.Now().Unix(),
	}
}
