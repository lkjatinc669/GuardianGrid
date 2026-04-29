package activeusers

import (
	"time"

	"github.com/shirou/gopsutil/v3/host"
)

// PUBLIC
func ScanActiveUsers() (map[string]interface{}, error) {
	users, err := host.Users()
	if err != nil {
		return emptyUsers(), nil
	}

	result := make([]map[string]interface{}, 0, len(users))
	seen := make(map[string]bool)

	for _, u := range users {
		if u.User == "" {
			continue
		}

		key := u.User + "|" + u.Terminal + "|" + u.Host
		if seen[key] {
			continue
		}
		seen[key] = true

		result = append(result, map[string]interface{}{
			"user":     u.User,
			"terminal": u.Terminal,
			"host":     u.Host,
			"started":  int64(u.Started), // unix timestamp
		})
	}

	return map[string]interface{}{
		"active_users": result,
		"count":        len(result),
		"timestamp":    time.Now().Unix(),
	}, nil
}
