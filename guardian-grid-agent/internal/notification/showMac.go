//go:build darwin

package notification

import (
	"fmt"
	"os/exec"
)

func showNotification(title, message string, nType NotificationType) {
	cmd := exec.Command(
		"osascript",
		"-e",
		fmt.Sprintf(
			`display notification "%s" with title "%s"`,
			message,
			title,
		),
	)

	cmd.Run()
}
