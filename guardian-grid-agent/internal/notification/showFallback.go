//go:build !windows && !linux && !darwin

package notification

import "fmt"

func showNotification(title, message string, nType NotificationType) {
	// Fallback to console print if no GUI notification system is available
	fmt.Printf("[%s] %s: %s\n", nType, title, message)
}
