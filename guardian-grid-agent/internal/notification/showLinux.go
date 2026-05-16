//go:build linux

package notification

import "os/exec"

func showLinux(title, message string, nType NotificationType) {
	urgency := "low"

	switch nType {
	case Error:
		urgency = "critical"
	case Warning:
		urgency = "normal"
	case Info:
		urgency = "low"
	}

	cmd := exec.Command("notify-send",
		"-u", urgency,
		title,
		message,
	)

	cmd.Run()
}
