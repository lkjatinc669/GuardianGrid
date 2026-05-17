//go:build linux

package notification

import "os/exec"

func showNotification(title, message string, nType NotificationType) {
	urgency := "low"

	switch nType {
	case Error:
		urgency = "critical"
	case Warning:
		urgency = "normal"
	}

	cmd := exec.Command(
		"notify-send",
		"-u", urgency,
		title,
		message,
	)

	cmd.Run()
}
